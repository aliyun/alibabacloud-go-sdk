// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ChangeResourceGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *ChangeResourceGroupResponseBody
	GetMessage() *string
	SetNewResourceGroupId(v string) *ChangeResourceGroupResponseBody
	GetNewResourceGroupId() *string
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
	SetSuccess(v int64) *ChangeResourceGroupResponseBody
	GetSuccess() *int64
}

type ChangeResourceGroupResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the new resource group. You can view the available resource groups in the Resource Management console.
	//
	// example:
	//
	// rg-ac***********7q
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C5CA600C-7D5A-45B5-B6DB-44FAC2C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *int64 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ChangeResourceGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChangeResourceGroupResponseBody) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) GetSuccess() *int64 {
	return s.Success
}

func (s *ChangeResourceGroupResponseBody) SetCode(v int32) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetNewResourceGroupId(v string) *ChangeResourceGroupResponseBody {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetSuccess(v int64) *ChangeResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
