// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseUserAllowOrgCreationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLangfuseUserAllowOrgCreationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLangfuseUserAllowOrgCreationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLangfuseUserAllowOrgCreationResponseBody) *ModifyLangfuseUserAllowOrgCreationResponse
	GetBody() *ModifyLangfuseUserAllowOrgCreationResponseBody
}

type ModifyLangfuseUserAllowOrgCreationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLangfuseUserAllowOrgCreationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLangfuseUserAllowOrgCreationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseUserAllowOrgCreationResponse) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseUserAllowOrgCreationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLangfuseUserAllowOrgCreationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLangfuseUserAllowOrgCreationResponse) GetBody() *ModifyLangfuseUserAllowOrgCreationResponseBody {
	return s.Body
}

func (s *ModifyLangfuseUserAllowOrgCreationResponse) SetHeaders(v map[string]*string) *ModifyLangfuseUserAllowOrgCreationResponse {
	s.Headers = v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationResponse) SetStatusCode(v int32) *ModifyLangfuseUserAllowOrgCreationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationResponse) SetBody(v *ModifyLangfuseUserAllowOrgCreationResponseBody) *ModifyLangfuseUserAllowOrgCreationResponse {
	s.Body = v
	return s
}

func (s *ModifyLangfuseUserAllowOrgCreationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
