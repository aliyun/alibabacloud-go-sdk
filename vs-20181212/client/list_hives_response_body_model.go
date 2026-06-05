// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHivesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHives(v []*ListHivesResponseBodyHives) *ListHivesResponseBody
	GetHives() []*ListHivesResponseBodyHives
	SetRequestId(v string) *ListHivesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListHivesResponseBody
	GetTotalCount() *int64
}

type ListHivesResponseBody struct {
	Hives []*ListHivesResponseBodyHives `json:"Hives,omitempty" xml:"Hives,omitempty" type:"Repeated"`
	// example:
	//
	// xxxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHivesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHivesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHivesResponseBody) GetHives() []*ListHivesResponseBodyHives {
	return s.Hives
}

func (s *ListHivesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHivesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListHivesResponseBody) SetHives(v []*ListHivesResponseBodyHives) *ListHivesResponseBody {
	s.Hives = v
	return s
}

func (s *ListHivesResponseBody) SetRequestId(v string) *ListHivesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHivesResponseBody) SetTotalCount(v int64) *ListHivesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHivesResponseBody) Validate() error {
	if s.Hives != nil {
		for _, item := range s.Hives {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHivesResponseBodyHives struct {
	// example:
	//
	// 2025-05-14T15:20:37+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hive-3b506f0868a7451ba15e0e890706033a
	HiveId *string `json:"HiveId,omitempty" xml:"HiveId,omitempty"`
	// example:
	//
	// yy-test2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListHivesResponseBodyHives) String() string {
	return dara.Prettify(s)
}

func (s ListHivesResponseBodyHives) GoString() string {
	return s.String()
}

func (s *ListHivesResponseBodyHives) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListHivesResponseBodyHives) GetDescription() *string {
	return s.Description
}

func (s *ListHivesResponseBodyHives) GetHiveId() *string {
	return s.HiveId
}

func (s *ListHivesResponseBodyHives) GetName() *string {
	return s.Name
}

func (s *ListHivesResponseBodyHives) SetCreationTime(v string) *ListHivesResponseBodyHives {
	s.CreationTime = &v
	return s
}

func (s *ListHivesResponseBodyHives) SetDescription(v string) *ListHivesResponseBodyHives {
	s.Description = &v
	return s
}

func (s *ListHivesResponseBodyHives) SetHiveId(v string) *ListHivesResponseBodyHives {
	s.HiveId = &v
	return s
}

func (s *ListHivesResponseBodyHives) SetName(v string) *ListHivesResponseBodyHives {
	s.Name = &v
	return s
}

func (s *ListHivesResponseBodyHives) Validate() error {
	return dara.Validate(s)
}
