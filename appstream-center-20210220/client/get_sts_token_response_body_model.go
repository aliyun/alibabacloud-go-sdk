// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStsTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStsTokenResponseBody
	GetRequestId() *string
	SetStsTokenModel(v *GetStsTokenResponseBodyStsTokenModel) *GetStsTokenResponseBody
	GetStsTokenModel() *GetStsTokenResponseBodyStsTokenModel
}

type GetStsTokenResponseBody struct {
	// example:
	//
	// CCF92035-6231-5ABB-930E-1E003C32****
	RequestId     *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StsTokenModel *GetStsTokenResponseBodyStsTokenModel `json:"StsTokenModel,omitempty" xml:"StsTokenModel,omitempty" type:"Struct"`
}

func (s GetStsTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStsTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetStsTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStsTokenResponseBody) GetStsTokenModel() *GetStsTokenResponseBodyStsTokenModel {
	return s.StsTokenModel
}

func (s *GetStsTokenResponseBody) SetRequestId(v string) *GetStsTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStsTokenResponseBody) SetStsTokenModel(v *GetStsTokenResponseBodyStsTokenModel) *GetStsTokenResponseBody {
	s.StsTokenModel = v
	return s
}

func (s *GetStsTokenResponseBody) Validate() error {
	if s.StsTokenModel != nil {
		if err := s.StsTokenModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStsTokenResponseBodyStsTokenModel struct {
	// example:
	//
	// be4be09e-cd00-4b4c-add7-11b4d8****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// sts****
	StsToken *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
	// example:
	//
	// 105552640689****
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetStsTokenResponseBodyStsTokenModel) String() string {
	return dara.Prettify(s)
}

func (s GetStsTokenResponseBodyStsTokenModel) GoString() string {
	return s.String()
}

func (s *GetStsTokenResponseBodyStsTokenModel) GetSessionId() *string {
	return s.SessionId
}

func (s *GetStsTokenResponseBodyStsTokenModel) GetStsToken() *string {
	return s.StsToken
}

func (s *GetStsTokenResponseBodyStsTokenModel) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetStsTokenResponseBodyStsTokenModel) SetSessionId(v string) *GetStsTokenResponseBodyStsTokenModel {
	s.SessionId = &v
	return s
}

func (s *GetStsTokenResponseBodyStsTokenModel) SetStsToken(v string) *GetStsTokenResponseBodyStsTokenModel {
	s.StsToken = &v
	return s
}

func (s *GetStsTokenResponseBodyStsTokenModel) SetTenantId(v int64) *GetStsTokenResponseBodyStsTokenModel {
	s.TenantId = &v
	return s
}

func (s *GetStsTokenResponseBodyStsTokenModel) Validate() error {
	return dara.Validate(s)
}
