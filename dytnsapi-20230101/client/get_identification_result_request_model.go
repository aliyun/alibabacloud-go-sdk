// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentificationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GetIdentificationResultRequest
	GetAuthCode() *string
	SetSessionId(v string) *GetIdentificationResultRequest
	GetSessionId() *string
}

type GetIdentificationResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetIdentificationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIdentificationResultRequest) GoString() string {
	return s.String()
}

func (s *GetIdentificationResultRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetIdentificationResultRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetIdentificationResultRequest) SetAuthCode(v string) *GetIdentificationResultRequest {
	s.AuthCode = &v
	return s
}

func (s *GetIdentificationResultRequest) SetSessionId(v string) *GetIdentificationResultRequest {
	s.SessionId = &v
	return s
}

func (s *GetIdentificationResultRequest) Validate() error {
	return dara.Validate(s)
}
