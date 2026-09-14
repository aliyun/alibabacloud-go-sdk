// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateComputeResourceAuthUserMappingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateComputeResourceAuthUserMappingsResponseBodyData) *UpdateComputeResourceAuthUserMappingsResponseBody
	GetData() *UpdateComputeResourceAuthUserMappingsResponseBodyData
	SetRequestId(v string) *UpdateComputeResourceAuthUserMappingsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateComputeResourceAuthUserMappingsResponseBody
	GetSuccess() *bool
}

type UpdateComputeResourceAuthUserMappingsResponseBody struct {
	// The data object.
	Data *UpdateComputeResourceAuthUserMappingsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 10000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateComputeResourceAuthUserMappingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceAuthUserMappingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBody) GetData() *UpdateComputeResourceAuthUserMappingsResponseBodyData {
	return s.Data
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBody) SetData(v *UpdateComputeResourceAuthUserMappingsResponseBodyData) *UpdateComputeResourceAuthUserMappingsResponseBody {
	s.Data = v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBody) SetRequestId(v string) *UpdateComputeResourceAuthUserMappingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBody) SetSuccess(v bool) *UpdateComputeResourceAuthUserMappingsResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateComputeResourceAuthUserMappingsResponseBodyData struct {
	// The change record ID.
	//
	// example:
	//
	// 123xx
	ChangeRecordId *int64 `json:"ChangeRecordId,omitempty" xml:"ChangeRecordId,omitempty"`
	// Indicates whether the operation succeeded. Valid values:
	//
	// - success: The update succeeded.
	//
	// - fail: The update failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateComputeResourceAuthUserMappingsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateComputeResourceAuthUserMappingsResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBodyData) GetChangeRecordId() *int64 {
	return s.ChangeRecordId
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBodyData) SetChangeRecordId(v int64) *UpdateComputeResourceAuthUserMappingsResponseBodyData {
	s.ChangeRecordId = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBodyData) SetStatus(v string) *UpdateComputeResourceAuthUserMappingsResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateComputeResourceAuthUserMappingsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
