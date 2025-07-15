// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRequestLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestLogId(v string) *QueryRequestLogsRequest
	GetRequestLogId() *string
	SetSecurityToken(v string) *QueryRequestLogsRequest
	GetSecurityToken() *string
}

type QueryRequestLogsRequest struct {
	// The ID of the request log.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95657ED9-2F6F-426F-BD99-79C8********
	RequestLogId  *string `json:"RequestLogId,omitempty" xml:"RequestLogId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s QueryRequestLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRequestLogsRequest) GoString() string {
	return s.String()
}

func (s *QueryRequestLogsRequest) GetRequestLogId() *string {
	return s.RequestLogId
}

func (s *QueryRequestLogsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *QueryRequestLogsRequest) SetRequestLogId(v string) *QueryRequestLogsRequest {
	s.RequestLogId = &v
	return s
}

func (s *QueryRequestLogsRequest) SetSecurityToken(v string) *QueryRequestLogsRequest {
	s.SecurityToken = &v
	return s
}

func (s *QueryRequestLogsRequest) Validate() error {
	return dara.Validate(s)
}
