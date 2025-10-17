// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivyComputeTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateLivyComputeTokenResponseBody
	GetCode() *string
	SetData(v *CreateLivyComputeTokenResponseBodyData) *CreateLivyComputeTokenResponseBody
	GetData() *CreateLivyComputeTokenResponseBodyData
	SetMessage(v string) *CreateLivyComputeTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateLivyComputeTokenResponseBody
	GetRequestId() *string
}

type CreateLivyComputeTokenResponseBody struct {
	// example:
	//
	// 1000000
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateLivyComputeTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateLivyComputeTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateLivyComputeTokenResponseBody) GetData() *CreateLivyComputeTokenResponseBodyData {
	return s.Data
}

func (s *CreateLivyComputeTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateLivyComputeTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLivyComputeTokenResponseBody) SetCode(v string) *CreateLivyComputeTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLivyComputeTokenResponseBody) SetData(v *CreateLivyComputeTokenResponseBodyData) *CreateLivyComputeTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateLivyComputeTokenResponseBody) SetMessage(v string) *CreateLivyComputeTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLivyComputeTokenResponseBody) SetRequestId(v string) *CreateLivyComputeTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLivyComputeTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateLivyComputeTokenResponseBodyData struct {
	// Token ID。
	//
	// example:
	//
	// lctk-xxxxxxxx
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s CreateLivyComputeTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateLivyComputeTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLivyComputeTokenResponseBodyData) GetTokenId() *string {
	return s.TokenId
}

func (s *CreateLivyComputeTokenResponseBodyData) SetTokenId(v string) *CreateLivyComputeTokenResponseBodyData {
	s.TokenId = &v
	return s
}

func (s *CreateLivyComputeTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
