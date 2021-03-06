/*
Real-time Online/Offline Charging System (OCS) for Telecom & ISP environments
Copyright (C) ITsysCOM GmbH

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/
package engine

import (
	"reflect"
	"testing"
	"time"

	"github.com/cgrates/cgrates/utils"
)

var rLS *ResourceLimiterService

/*
func TestRLsIndexStringFilters(t *testing.T) {
	rls := []*ResourceLimit{
		&ResourceLimit{
			ID:     "RL1",
			Weight: 20,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"1001", "1002"}},
				&RequestFilter{Type: MetaRSRFields, Values: []string{"Subject(~^1.*1$)", "Destination(1002)"},
					rsrFields: utils.ParseRSRFieldsMustCompile("Subject(~^1.*1$);Destination(1002)", utils.INFIELD_SEP),
				}},
			ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC),
			Limit:          2,
			Usage:          make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL2",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"dan", "1002"}},
				&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan"}},
			},
			ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC),
			Limit:          1,
			UsageTTL:       time.Duration(1 * time.Millisecond),
			Usage:          make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL4",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaStringPrefix, FieldName: "Destination", Values: []string{"+49"}},
			},
			ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC),
			Limit:          1,
			Usage:          make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL5",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaStringPrefix, FieldName: "Destination", Values: []string{"+40"}},
			},
			ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC),
			Limit:          1,
			UsageTTL:       time.Duration(10 * time.Millisecond),
			Usage:          make(map[string]*ResourceUsage),
		},
	}
	for _, rl := range rls {
		cache.Set(utils.ResourceLimitsPrefix+rl.ID, rl, true, "")
	}
	rLS = new(ResourceLimiterService)
	eIndexes := map[string]map[string]utils.StringMap{
		"Account": map[string]utils.StringMap{
			"1001": utils.StringMap{
				"RL1": true,
			},
			"1002": utils.StringMap{
				"RL1": true,
				"RL2": true,
			},
			"dan": utils.StringMap{
				"RL2": true,
			},
		},
		"Subject": map[string]utils.StringMap{
			"dan": utils.StringMap{
				"RL2": true,
			},
		},
		utils.NOT_AVAILABLE: map[string]utils.StringMap{
			utils.NOT_AVAILABLE: utils.StringMap{
				"RL4": true,
				"RL5": true,
			},
		},
	}
	if err := rLS.indexStringFilters(nil); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(eIndexes, rLS.stringIndexes) {
		t.Errorf("Expecting: %+v, received: %+v", eIndexes, rLS.stringIndexes)
	}
	rl3 := &ResourceLimit{
		ID:     "RL3",
		Weight: 10,
		Filters: []*RequestFilter{
			&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan"}},
			&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"1003"}},
		},
		ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC),
		Limit:          1,
		Usage:          make(map[string]*ResourceUsage),
	}
	cache.Set(utils.ResourceLimitsPrefix+rl3.ID, rl3, true, "")
	rl6 := &ResourceLimit{ // Add it so we can test expiryTime
		ID:     "RL6",
		Weight: 10,
		Filters: []*RequestFilter{
			&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan"}},
		},
		ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC),
		ExpiryTime:     time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC),
		Limit:          1,
		Usage:          make(map[string]*ResourceUsage),
	}
	cache.Set(utils.ResourceLimitsPrefix+rl6.ID, rl6, true, "")
	eIndexes = map[string]map[string]utils.StringMap{
		"Account": map[string]utils.StringMap{
			"1001": utils.StringMap{
				"RL1": true,
			},
			"1002": utils.StringMap{
				"RL1": true,
				"RL2": true,
			},
			"dan": utils.StringMap{
				"RL2": true,
			},
		},
		"Subject": map[string]utils.StringMap{
			"dan": utils.StringMap{
				"RL2": true,
				"RL3": true,
				"RL6": true,
			},
			"1003": utils.StringMap{
				"RL3": true,
			},
		},
		utils.NOT_AVAILABLE: map[string]utils.StringMap{
			utils.NOT_AVAILABLE: utils.StringMap{
				"RL4": true,
				"RL5": true,
			},
		},
	}
	// Test index update
	if err := rLS.indexStringFilters([]string{rl3.ID, rl6.ID}); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(eIndexes, rLS.stringIndexes) {
		t.Errorf("Expecting: %+v, received: %+v", eIndexes, rLS.stringIndexes)
	}
}
*/
/*
func TestResourceLimitsAllowUsage(t *testing.T) {
	rls := make(ResourceLimits, 0)
	if !rls.AllowUsage(1.0) {
		t.Error("Not allowed for empty limits")
	}
	rls = ResourceLimits{
		&ResourceLimit{
			ID:     "RLAU1",
			Weight: 20,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"1001", "1002"}},
				&RequestFilter{Type: MetaRSRFields, Values: []string{"Subject(~^1.*1$)", "Destination(1002)"},
					rsrFields: utils.ParseRSRFieldsMustCompile("Subject(~^1.*1$);Destination(1002)", utils.INFIELD_SEP),
				}},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              1,
			Usage: map[string]*ResourceUsage{"call1": &ResourceUsage{
				ID: "call1", UsageTime: time.Now(), UsageUnits: 1}},
			TotalUsage: 1,
		},
		&ResourceLimit{
			ID:     "RLAU2",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"dan", "1002"}},
				&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan"}},
			},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              2,
			UsageTTL:           time.Duration(1 * time.Millisecond),
			Usage:              make(map[string]*ResourceUsage),
		},
	}
	if !rls.AllowUsage(2.0) {
		t.Error("Not allowed")
	}
}
*/

func TestRLsLoadRLs(t *testing.T) {
	rls := []*ResourceLimit{
		&ResourceLimit{
			ID:     "RL1",
			Weight: 20,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"1001", "1002"}},
				&RequestFilter{Type: MetaRSRFields, Values: []string{"Subject(~^1.*1$)", "Destination(1002)"},
					rsrFields: utils.ParseRSRFieldsMustCompile("Subject(~^1.*1$);Destination(1002)", utils.INFIELD_SEP),
				}},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              2,
			Usage:              make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL2",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"dan", "1002"}},
				&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan"}},
			},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              1,
			UsageTTL:           time.Duration(1 * time.Millisecond),
			Usage:              make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL3",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan2"}},
				&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"1003"}},
			},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              1,
			Usage:              make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL4",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaStringPrefix, FieldName: "Destination", Values: []string{"+49"}},
			},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              1,
			Usage:              make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL5",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaStringPrefix, FieldName: "Destination", Values: []string{"+40"}},
			},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              1,
			UsageTTL:           time.Duration(10 * time.Millisecond),
			Usage:              make(map[string]*ResourceUsage),
		},
		&ResourceLimit{ // Add it so we can test expiryTime
			ID:     "RL6",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan2"}},
			},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			ExpiryTime:         time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC),
			Limit:              1,
			Usage:              make(map[string]*ResourceUsage),
		},
	}
	rlIdxr, err := NewReqFilterIndexer(dataStorage, utils.ResourceLimitsIndex)
	if err != nil {
		t.Error(err)
	}
	for _, rl := range rls {
		if err := dataStorage.SetResourceLimit(rl, utils.NonTransactional); err != nil {
			t.Error(err)
		}
		rlIdxr.IndexFilters(rl.ID, rl.Filters)
	}
	if err := rlIdxr.StoreIndexes(); err != nil {
		t.Error(err)
	}
}

func TestRLsMatchingResourceLimitsForEvent(t *testing.T) {
	rLS = &ResourceLimiterService{dataDB: dataStorage, cdrStatS: nil}
	eResLimits := ResourceLimits{
		&ResourceLimit{
			ID:     "RL1",
			Weight: 20,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"1001", "1002"}},
				&RequestFilter{Type: MetaRSRFields, Values: []string{"Subject(~^1.*1$)", "Destination(1002)"},
					rsrFields: utils.ParseRSRFieldsMustCompile("Subject(~^1.*1$);Destination(1002)", utils.INFIELD_SEP),
				}},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              2,
			Usage:              make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL2",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"dan", "1002"}},
				&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan"}},
			},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              1,
			UsageTTL:           time.Duration(1 * time.Millisecond),
			Usage:              make(map[string]*ResourceUsage),
		},
	}
	if resLimits, err := rLS.matchingResourceLimitsForEvent(
		map[string]interface{}{"Account": "1002", "Subject": "dan", "Destination": "1002"}); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(eResLimits[0].Filters[0], resLimits[0].Filters[0]) {
		t.Errorf("Expecting: %+v, received: %+v", eResLimits, resLimits)
	} else {
		// Make sure the filters are what we expect to be after retrieving from cache:
		fltr := resLimits[0].Filters[1]
		if pass, _ := fltr.Pass(map[string]interface{}{"Subject": "10000001"}, "", nil); !pass {
			t.Errorf("Expecting RL: %+v, received: %+v", eResLimits[0], resLimits[0])
		}
		if pass, _ := fltr.Pass(map[string]interface{}{"Account": "1002"}, "", nil); pass {
			t.Errorf("Expecting RL: %+v, received: %+v", eResLimits[0], resLimits[0])
		}
	}
}

func TestRLsV1ResourceLimitsForEvent(t *testing.T) {
	eLimits := []*ResourceLimit{
		&ResourceLimit{
			ID:     "RL1",
			Weight: 20,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"1001", "1002"}},
				&RequestFilter{Type: MetaRSRFields, Values: []string{"Subject(~^1.*1$)", "Destination(1002)"},
					rsrFields: utils.ParseRSRFieldsMustCompile("Subject(~^1.*1$);Destination(1002)", utils.INFIELD_SEP),
				}},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              2,
			Usage:              make(map[string]*ResourceUsage),
		},
		&ResourceLimit{
			ID:     "RL2",
			Weight: 10,
			Filters: []*RequestFilter{
				&RequestFilter{Type: MetaString, FieldName: "Account", Values: []string{"dan", "1002"}},
				&RequestFilter{Type: MetaString, FieldName: "Subject", Values: []string{"dan"}},
			},
			ActivationInterval: &utils.ActivationInterval{ActivationTime: time.Date(2014, 7, 3, 13, 43, 0, 1, time.UTC)},
			Limit:              1,
			UsageTTL:           time.Duration(1 * time.Millisecond),
			Usage:              make(map[string]*ResourceUsage),
		},
	}
	var rcvLmts []*ResourceLimit
	if err := rLS.V1ResourceLimitsForEvent(map[string]interface{}{"Account": "1002", "Subject": "dan", "Destination": "1002"}, &rcvLmts); err != nil {
		t.Error(err)
	} else if len(eLimits) != len(rcvLmts) {
		t.Errorf("Expecting: %+v, received: %+v", eLimits, rcvLmts)
	}
}

func TestRLsV1InitiateResourceUsage(t *testing.T) {
	attrRU := utils.AttrRLsResourceUsage{
		UsageID: "651a8db2-4f67-4cf8-b622-169e8a482e50",
		Event:   map[string]interface{}{"Account": "1002", "Subject": "dan", "Destination": "1002"},
		Units:   1,
	}
	var reply string
	if err := rLS.V1AllocateResource(attrRU, &reply); err != nil {
		t.Error(err)
	} else if reply != "RL1" {
		t.Error("Received reply: ", reply)
	}
	resLimits, err := rLS.matchingResourceLimitsForEvent(attrRU.Event)
	if err != nil {
		t.Error(err)
	} else if len(resLimits) != 2 {
		t.Errorf("Received: %+v", resLimits)
	} else if resLimits[0].UsedUnits() != 1 {
		t.Errorf("RL1: %+v", resLimits[0])
	} else if _, hasKey := resLimits[0].Usage[attrRU.UsageID]; !hasKey {
		t.Errorf("RL1: %+v", resLimits[0])
	}
}

func TestRLsV1TerminateResourceUsage(t *testing.T) {
	attrRU := utils.AttrRLsResourceUsage{
		UsageID: "651a8db2-4f67-4cf8-b622-169e8a482e50",
		Event:   map[string]interface{}{"Account": "1002", "Subject": "dan", "Destination": "1002"},
		Units:   1,
	}
	var reply string
	if err := rLS.V1ReleaseResource(attrRU, &reply); err != nil {
		t.Error(err)
	} else if reply != utils.OK {
		t.Error("Received reply: ", reply)
	}
	resLimits, err := rLS.matchingResourceLimitsForEvent(attrRU.Event)
	if err != nil {
		t.Error(err)
	} else if len(resLimits) != 2 {
		t.Errorf("Received: %+v", resLimits)
	} else if resLimits[0].UsedUnits() != 0 {
		t.Errorf("RL1: %+v", resLimits[0])
	} else if _, hasKey := resLimits[0].Usage[attrRU.UsageID]; hasKey {
		t.Errorf("RL1: %+v", resLimits[0])
	}
}
