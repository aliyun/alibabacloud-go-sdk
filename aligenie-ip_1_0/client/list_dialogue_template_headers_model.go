// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDialogueTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDialogueTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetXAcsAligenieAccessToken(v string) *ListDialogueTemplateHeaders
	GetXAcsAligenieAccessToken() *string
	SetAuthorization(v string) *ListDialogueTemplateHeaders
	GetAuthorization() *string
}

type ListDialogueTemplateHeaders struct {
	CommonHeaders           map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	XAcsAligenieAccessToken *string            `json:"x-acs-aligenie-access-token,omitempty" xml:"x-acs-aligenie-access-token,omitempty"`
	Authorization           *string            `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ListDialogueTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDialogueTemplateHeaders) GoString() string {
	return s.String()
}

func (s *ListDialogueTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDialogueTemplateHeaders) GetXAcsAligenieAccessToken() *string {
	return s.XAcsAligenieAccessToken
}

func (s *ListDialogueTemplateHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ListDialogueTemplateHeaders) SetCommonHeaders(v map[string]*string) *ListDialogueTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDialogueTemplateHeaders) SetXAcsAligenieAccessToken(v string) *ListDialogueTemplateHeaders {
	s.XAcsAligenieAccessToken = &v
	return s
}

func (s *ListDialogueTemplateHeaders) SetAuthorization(v string) *ListDialogueTemplateHeaders {
	s.Authorization = &v
	return s
}

func (s *ListDialogueTemplateHeaders) Validate() error {
	return dara.Validate(s)
}
