// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMetaEntityResponseBody
	GetRequestId() *string
	SetResult(v *UpdateMetaEntityResponseBodyResult) *UpdateMetaEntityResponseBody
	GetResult() *UpdateMetaEntityResponseBodyResult
	SetSuccess(v bool) *UpdateMetaEntityResponseBody
	GetSuccess() *bool
}

type UpdateMetaEntityResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// AASFDFSDFG-DFSDF-DFSDFD-SDFSDF
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *UpdateMetaEntityResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMetaEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaEntityResponseBody) GetResult() *UpdateMetaEntityResponseBodyResult {
	return s.Result
}

func (s *UpdateMetaEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMetaEntityResponseBody) SetRequestId(v string) *UpdateMetaEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaEntityResponseBody) SetResult(v *UpdateMetaEntityResponseBodyResult) *UpdateMetaEntityResponseBody {
	s.Result = v
	return s
}

func (s *UpdateMetaEntityResponseBody) SetSuccess(v bool) *UpdateMetaEntityResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMetaEntityResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMetaEntityResponseBodyResult struct {
	// example:
	//
	// custom_entity-customer_api:api_001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMetaEntityResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *UpdateMetaEntityResponseBodyResult) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMetaEntityResponseBodyResult) SetId(v string) *UpdateMetaEntityResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateMetaEntityResponseBodyResult) SetSuccess(v bool) *UpdateMetaEntityResponseBodyResult {
	s.Success = &v
	return s
}

func (s *UpdateMetaEntityResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
