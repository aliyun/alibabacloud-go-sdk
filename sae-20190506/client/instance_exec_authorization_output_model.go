// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceExecAuthorizationOutput interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKeyId(v string) *InstanceExecAuthorizationOutput
	GetAccessKeyId() *string
	SetAccountId(v string) *InstanceExecAuthorizationOutput
	GetAccountId() *string
	SetAuthorization(v string) *InstanceExecAuthorizationOutput
	GetAuthorization() *string
	SetDate(v string) *InstanceExecAuthorizationOutput
	GetDate() *string
	SetEndpoint(v string) *InstanceExecAuthorizationOutput
	GetEndpoint() *string
	SetRequestId(v string) *InstanceExecAuthorizationOutput
	GetRequestId() *string
}

type InstanceExecAuthorizationOutput struct {
	AccessKeyId   *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	AccountId     *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Authorization *string `json:"authorization,omitempty" xml:"authorization,omitempty"`
	Date          *string `json:"date,omitempty" xml:"date,omitempty"`
	Endpoint      *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	RequestId     *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s InstanceExecAuthorizationOutput) String() string {
	return dara.Prettify(s)
}

func (s InstanceExecAuthorizationOutput) GoString() string {
	return s.String()
}

func (s *InstanceExecAuthorizationOutput) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *InstanceExecAuthorizationOutput) GetAccountId() *string {
	return s.AccountId
}

func (s *InstanceExecAuthorizationOutput) GetAuthorization() *string {
	return s.Authorization
}

func (s *InstanceExecAuthorizationOutput) GetDate() *string {
	return s.Date
}

func (s *InstanceExecAuthorizationOutput) GetEndpoint() *string {
	return s.Endpoint
}

func (s *InstanceExecAuthorizationOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *InstanceExecAuthorizationOutput) SetAccessKeyId(v string) *InstanceExecAuthorizationOutput {
	s.AccessKeyId = &v
	return s
}

func (s *InstanceExecAuthorizationOutput) SetAccountId(v string) *InstanceExecAuthorizationOutput {
	s.AccountId = &v
	return s
}

func (s *InstanceExecAuthorizationOutput) SetAuthorization(v string) *InstanceExecAuthorizationOutput {
	s.Authorization = &v
	return s
}

func (s *InstanceExecAuthorizationOutput) SetDate(v string) *InstanceExecAuthorizationOutput {
	s.Date = &v
	return s
}

func (s *InstanceExecAuthorizationOutput) SetEndpoint(v string) *InstanceExecAuthorizationOutput {
	s.Endpoint = &v
	return s
}

func (s *InstanceExecAuthorizationOutput) SetRequestId(v string) *InstanceExecAuthorizationOutput {
	s.RequestId = &v
	return s
}

func (s *InstanceExecAuthorizationOutput) Validate() error {
	return dara.Validate(s)
}
