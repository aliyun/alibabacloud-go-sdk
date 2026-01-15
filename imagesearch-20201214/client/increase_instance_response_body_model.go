// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncreaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *IncreaseInstanceResponseBodyData) *IncreaseInstanceResponseBody
	GetData() *IncreaseInstanceResponseBodyData
	SetRequestId(v string) *IncreaseInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IncreaseInstanceResponseBody
	GetSuccess() *bool
}

type IncreaseInstanceResponseBody struct {
	// The information about the task.
	Data *IncreaseInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 36C43E96-8F68-44AA-B1AF-B1F7AB94A6C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IncreaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IncreaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponseBody) GetData() *IncreaseInstanceResponseBodyData {
	return s.Data
}

func (s *IncreaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IncreaseInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IncreaseInstanceResponseBody) SetData(v *IncreaseInstanceResponseBodyData) *IncreaseInstanceResponseBody {
	s.Data = v
	return s
}

func (s *IncreaseInstanceResponseBody) SetRequestId(v string) *IncreaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *IncreaseInstanceResponseBody) SetSuccess(v bool) *IncreaseInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *IncreaseInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IncreaseInstanceResponseBodyData struct {
	// The ID of the task.
	//
	// example:
	//
	// 500
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The status of the task.
	//
	// 	- PROCESSING: in progress
	//
	// 	- FAIL: failed
	//
	// 	- SUCCESS: successful
	//
	// example:
	//
	// PROCESSING
	IncrementStatus *string `json:"IncrementStatus,omitempty" xml:"IncrementStatus,omitempty"`
}

func (s IncreaseInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s IncreaseInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *IncreaseInstanceResponseBodyData) GetId() *string {
	return s.Id
}

func (s *IncreaseInstanceResponseBodyData) GetIncrementStatus() *string {
	return s.IncrementStatus
}

func (s *IncreaseInstanceResponseBodyData) SetId(v string) *IncreaseInstanceResponseBodyData {
	s.Id = &v
	return s
}

func (s *IncreaseInstanceResponseBodyData) SetIncrementStatus(v string) *IncreaseInstanceResponseBodyData {
	s.IncrementStatus = &v
	return s
}

func (s *IncreaseInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
