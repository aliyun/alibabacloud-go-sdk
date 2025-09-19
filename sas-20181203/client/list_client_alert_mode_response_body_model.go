// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientAlertModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListClientAlertModeResponseBodyData) *ListClientAlertModeResponseBody
	GetData() *ListClientAlertModeResponseBodyData
	SetRequestId(v string) *ListClientAlertModeResponseBody
	GetRequestId() *string
}

type ListClientAlertModeResponseBody struct {
	// The response parameters.
	Data *ListClientAlertModeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientAlertModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientAlertModeResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientAlertModeResponseBody) GetData() *ListClientAlertModeResponseBodyData {
	return s.Data
}

func (s *ListClientAlertModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientAlertModeResponseBody) SetData(v *ListClientAlertModeResponseBodyData) *ListClientAlertModeResponseBody {
	s.Data = v
	return s
}

func (s *ListClientAlertModeResponseBody) SetRequestId(v string) *ListClientAlertModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientAlertModeResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClientAlertModeResponseBodyData struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The UUIDs of the assets.
	Uuids []*string `json:"Uuids,omitempty" xml:"Uuids,omitempty" type:"Repeated"`
}

func (s ListClientAlertModeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClientAlertModeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClientAlertModeResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *ListClientAlertModeResponseBodyData) GetUuids() []*string {
	return s.Uuids
}

func (s *ListClientAlertModeResponseBodyData) SetCount(v int32) *ListClientAlertModeResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListClientAlertModeResponseBodyData) SetUuids(v []*string) *ListClientAlertModeResponseBodyData {
	s.Uuids = v
	return s
}

func (s *ListClientAlertModeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
