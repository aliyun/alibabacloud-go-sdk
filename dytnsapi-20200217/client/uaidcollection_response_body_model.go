// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UAIDCollectionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UAIDCollectionResponseBody
	GetCode() *string
	SetMessage(v string) *UAIDCollectionResponseBody
	GetMessage() *string
	SetModel(v *UAIDCollectionResponseBodyModel) *UAIDCollectionResponseBody
	GetModel() *UAIDCollectionResponseBodyModel
	SetRequestId(v string) *UAIDCollectionResponseBody
	GetRequestId() *string
}

type UAIDCollectionResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *UAIDCollectionResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UAIDCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UAIDCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *UAIDCollectionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UAIDCollectionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UAIDCollectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UAIDCollectionResponseBody) GetModel() *UAIDCollectionResponseBodyModel {
	return s.Model
}

func (s *UAIDCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UAIDCollectionResponseBody) SetAccessDeniedDetail(v string) *UAIDCollectionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UAIDCollectionResponseBody) SetCode(v string) *UAIDCollectionResponseBody {
	s.Code = &v
	return s
}

func (s *UAIDCollectionResponseBody) SetMessage(v string) *UAIDCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *UAIDCollectionResponseBody) SetModel(v *UAIDCollectionResponseBodyModel) *UAIDCollectionResponseBody {
	s.Model = v
	return s
}

func (s *UAIDCollectionResponseBody) SetRequestId(v string) *UAIDCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UAIDCollectionResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UAIDCollectionResponseBodyModel struct {
	// example:
	//
	// 示例值
	Uaid *string `json:"Uaid,omitempty" xml:"Uaid,omitempty"`
}

func (s UAIDCollectionResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s UAIDCollectionResponseBodyModel) GoString() string {
	return s.String()
}

func (s *UAIDCollectionResponseBodyModel) GetUaid() *string {
	return s.Uaid
}

func (s *UAIDCollectionResponseBodyModel) SetUaid(v string) *UAIDCollectionResponseBodyModel {
	s.Uaid = &v
	return s
}

func (s *UAIDCollectionResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
