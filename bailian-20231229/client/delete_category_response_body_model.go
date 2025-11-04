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
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// data
	Data *DeleteCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// workspace id is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
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
