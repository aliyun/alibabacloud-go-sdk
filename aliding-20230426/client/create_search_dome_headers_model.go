// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSearchDomeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateSearchDomeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateSearchDomeHeadersAccountContext) *CreateSearchDomeHeaders
	GetAccountContext() *CreateSearchDomeHeadersAccountContext
}

type CreateSearchDomeHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateSearchDomeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateSearchDomeHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchDomeHeaders) GoString() string {
	return s.String()
}

func (s *CreateSearchDomeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateSearchDomeHeaders) GetAccountContext() *CreateSearchDomeHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateSearchDomeHeaders) SetCommonHeaders(v map[string]*string) *CreateSearchDomeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateSearchDomeHeaders) SetAccountContext(v *CreateSearchDomeHeadersAccountContext) *CreateSearchDomeHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateSearchDomeHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSearchDomeHeadersAccountContext struct {
	// example:
	//
	// ba3a9b612345678d8fedf544ef69d19e
	UserToken *string `json:"userToken,omitempty" xml:"userToken,omitempty"`
}

func (s CreateSearchDomeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateSearchDomeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateSearchDomeHeadersAccountContext) GetUserToken() *string {
	return s.UserToken
}

func (s *CreateSearchDomeHeadersAccountContext) SetUserToken(v string) *CreateSearchDomeHeadersAccountContext {
	s.UserToken = &v
	return s
}

func (s *CreateSearchDomeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
