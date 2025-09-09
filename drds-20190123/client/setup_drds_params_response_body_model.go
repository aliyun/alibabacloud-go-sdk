// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupDrdsParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetupDrdsParamsResponseBody
	GetData() *bool
	SetRequestId(v string) *SetupDrdsParamsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetupDrdsParamsResponseBody
	GetSuccess() *bool
}

type SetupDrdsParamsResponseBody struct {
	// The returned results.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupDrdsParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetupDrdsParamsResponseBody) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetupDrdsParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetupDrdsParamsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetupDrdsParamsResponseBody) SetData(v bool) *SetupDrdsParamsResponseBody {
	s.Data = &v
	return s
}

func (s *SetupDrdsParamsResponseBody) SetRequestId(v string) *SetupDrdsParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupDrdsParamsResponseBody) SetSuccess(v bool) *SetupDrdsParamsResponseBody {
	s.Success = &v
	return s
}

func (s *SetupDrdsParamsResponseBody) Validate() error {
	return dara.Validate(s)
}
