// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningTasksHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetRunningTasksHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetRunningTasksHeadersAccountContext) *GetRunningTasksHeaders
	GetAccountContext() *GetRunningTasksHeadersAccountContext
}

type GetRunningTasksHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetRunningTasksHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetRunningTasksHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetRunningTasksHeaders) GoString() string {
	return s.String()
}

func (s *GetRunningTasksHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetRunningTasksHeaders) GetAccountContext() *GetRunningTasksHeadersAccountContext {
	return s.AccountContext
}

func (s *GetRunningTasksHeaders) SetCommonHeaders(v map[string]*string) *GetRunningTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRunningTasksHeaders) SetAccountContext(v *GetRunningTasksHeadersAccountContext) *GetRunningTasksHeaders {
	s.AccountContext = v
	return s
}

func (s *GetRunningTasksHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRunningTasksHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetRunningTasksHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetRunningTasksHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetRunningTasksHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetRunningTasksHeadersAccountContext) SetAccountId(v string) *GetRunningTasksHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetRunningTasksHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
