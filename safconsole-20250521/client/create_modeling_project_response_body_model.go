// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateModelingProjectResponseBody
	GetCode() *int64
	SetRequestId(v string) *CreateModelingProjectResponseBody
	GetRequestId() *string
	SetResultObject(v *CreateModelingProjectResponseBodyResultObject) *CreateModelingProjectResponseBody
	GetResultObject() *CreateModelingProjectResponseBodyResultObject
	SetSuccess(v bool) *CreateModelingProjectResponseBody
	GetSuccess() *bool
}

type CreateModelingProjectResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	ResultObject *CreateModelingProjectResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateModelingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelingProjectResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateModelingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelingProjectResponseBody) GetResultObject() *CreateModelingProjectResponseBodyResultObject {
	return s.ResultObject
}

func (s *CreateModelingProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateModelingProjectResponseBody) SetCode(v int64) *CreateModelingProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateModelingProjectResponseBody) SetRequestId(v string) *CreateModelingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelingProjectResponseBody) SetResultObject(v *CreateModelingProjectResponseBodyResultObject) *CreateModelingProjectResponseBody {
	s.ResultObject = v
	return s
}

func (s *CreateModelingProjectResponseBody) SetSuccess(v bool) *CreateModelingProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateModelingProjectResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateModelingProjectResponseBodyResultObject struct {
	// Project ID
	//
	// example:
	//
	// 01
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateModelingProjectResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s CreateModelingProjectResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *CreateModelingProjectResponseBodyResultObject) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateModelingProjectResponseBodyResultObject) SetProjectId(v string) *CreateModelingProjectResponseBodyResultObject {
	s.ProjectId = &v
	return s
}

func (s *CreateModelingProjectResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
