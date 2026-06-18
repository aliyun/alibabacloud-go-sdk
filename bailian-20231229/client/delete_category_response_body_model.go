// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteCategoryResponseBody
	GetCode() *string
	SetData(v *DeleteCategoryResponseBodyData) *DeleteCategoryResponseBody
	GetData() *DeleteCategoryResponseBodyData
	SetMessage(v string) *DeleteCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCategoryResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteCategoryResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DeleteCategoryResponseBody
	GetSuccess() *bool
}

type DeleteCategoryResponseBody struct {
	// The error code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data field returned by the operation.
	Data *DeleteCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// workspace id is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code returned by the operation.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteCategoryResponseBody) GetData() *DeleteCategoryResponseBodyData {
	return s.Data
}

func (s *DeleteCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCategoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCategoryResponseBody) SetCode(v string) *DeleteCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetData(v *DeleteCategoryResponseBodyData) *DeleteCategoryResponseBody {
	s.Data = v
	return s
}

func (s *DeleteCategoryResponseBody) SetMessage(v string) *DeleteCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetRequestId(v string) *DeleteCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetStatus(v string) *DeleteCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteCategoryResponseBody) SetSuccess(v bool) *DeleteCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteCategoryResponseBodyData struct {
	// The ID of the deleted category.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
}

func (s DeleteCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteCategoryResponseBodyData) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DeleteCategoryResponseBodyData) SetCategoryId(v string) *DeleteCategoryResponseBodyData {
	s.CategoryId = &v
	return s
}

func (s *DeleteCategoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
