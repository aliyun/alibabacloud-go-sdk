// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateQuotaResponseBodyData) *CreateQuotaResponseBody
	GetData() *CreateQuotaResponseBodyData
	SetRequestId(v string) *CreateQuotaResponseBody
	GetRequestId() *string
}

type CreateQuotaResponseBody struct {
	// Response parameters.
	Data *CreateQuotaResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 0bc520ad17171208978521777d742c
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaResponseBody) GetData() *CreateQuotaResponseBodyData {
	return s.Data
}

func (s *CreateQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQuotaResponseBody) SetData(v *CreateQuotaResponseBodyData) *CreateQuotaResponseBody {
	s.Data = v
	return s
}

func (s *CreateQuotaResponseBody) SetRequestId(v string) *CreateQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQuotaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateQuotaResponseBodyData struct {
	// Quota alias.
	//
	// example:
	//
	// os_PayAsYouGoQuota_p
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
}

func (s CreateQuotaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateQuotaResponseBodyData) GetNickName() *string {
	return s.NickName
}

func (s *CreateQuotaResponseBodyData) SetNickName(v string) *CreateQuotaResponseBodyData {
	s.NickName = &v
	return s
}

func (s *CreateQuotaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
