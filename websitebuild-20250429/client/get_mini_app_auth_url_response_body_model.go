// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppAuthUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMiniAppAuthUrlResponseBodyData) *GetMiniAppAuthUrlResponseBody
	GetData() *GetMiniAppAuthUrlResponseBodyData
	SetRequestId(v string) *GetMiniAppAuthUrlResponseBody
	GetRequestId() *string
}

type GetMiniAppAuthUrlResponseBody struct {
	// Result of the request.
	Data *GetMiniAppAuthUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// request ID
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMiniAppAuthUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppAuthUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMiniAppAuthUrlResponseBody) GetData() *GetMiniAppAuthUrlResponseBodyData {
	return s.Data
}

func (s *GetMiniAppAuthUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMiniAppAuthUrlResponseBody) SetData(v *GetMiniAppAuthUrlResponseBodyData) *GetMiniAppAuthUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetMiniAppAuthUrlResponseBody) SetRequestId(v string) *GetMiniAppAuthUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMiniAppAuthUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMiniAppAuthUrlResponseBodyData struct {
	// authorized address
	//
	// example:
	//
	// https://nschiper.oneclick.accounts.logi.com/identity/oauth2/token
	AuthUrl *string `json:"AuthUrl,omitempty" xml:"AuthUrl,omitempty"`
}

func (s GetMiniAppAuthUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppAuthUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMiniAppAuthUrlResponseBodyData) GetAuthUrl() *string {
	return s.AuthUrl
}

func (s *GetMiniAppAuthUrlResponseBodyData) SetAuthUrl(v string) *GetMiniAppAuthUrlResponseBodyData {
	s.AuthUrl = &v
	return s
}

func (s *GetMiniAppAuthUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
