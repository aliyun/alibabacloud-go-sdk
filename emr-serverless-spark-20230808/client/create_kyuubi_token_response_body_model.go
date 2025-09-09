// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKyuubiTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateKyuubiTokenResponseBodyData) *CreateKyuubiTokenResponseBody
	GetData() *CreateKyuubiTokenResponseBodyData
	SetRequestId(v string) *CreateKyuubiTokenResponseBody
	GetRequestId() *string
}

type CreateKyuubiTokenResponseBody struct {
	Data *CreateKyuubiTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateKyuubiTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKyuubiTokenResponseBody) GetData() *CreateKyuubiTokenResponseBodyData {
	return s.Data
}

func (s *CreateKyuubiTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKyuubiTokenResponseBody) SetData(v *CreateKyuubiTokenResponseBodyData) *CreateKyuubiTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateKyuubiTokenResponseBody) SetRequestId(v string) *CreateKyuubiTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKyuubiTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateKyuubiTokenResponseBodyData struct {
	// Token ID。
	//
	// example:
	//
	// tk-zpi0*****hdv4y
	TokenId *string `json:"tokenId,omitempty" xml:"tokenId,omitempty"`
}

func (s CreateKyuubiTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateKyuubiTokenResponseBodyData) GetTokenId() *string {
	return s.TokenId
}

func (s *CreateKyuubiTokenResponseBodyData) SetTokenId(v string) *CreateKyuubiTokenResponseBodyData {
	s.TokenId = &v
	return s
}

func (s *CreateKyuubiTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
