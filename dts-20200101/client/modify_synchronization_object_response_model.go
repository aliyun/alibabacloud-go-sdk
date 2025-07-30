// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySynchronizationObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySynchronizationObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySynchronizationObjectResponse
	GetStatusCode() *int32
	SetBody(v *ModifySynchronizationObjectResponseBody) *ModifySynchronizationObjectResponse
	GetBody() *ModifySynchronizationObjectResponseBody
}

type ModifySynchronizationObjectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySynchronizationObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySynchronizationObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySynchronizationObjectResponse) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySynchronizationObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySynchronizationObjectResponse) GetBody() *ModifySynchronizationObjectResponseBody {
	return s.Body
}

func (s *ModifySynchronizationObjectResponse) SetHeaders(v map[string]*string) *ModifySynchronizationObjectResponse {
	s.Headers = v
	return s
}

func (s *ModifySynchronizationObjectResponse) SetStatusCode(v int32) *ModifySynchronizationObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySynchronizationObjectResponse) SetBody(v *ModifySynchronizationObjectResponseBody) *ModifySynchronizationObjectResponse {
	s.Body = v
	return s
}

func (s *ModifySynchronizationObjectResponse) Validate() error {
	return dara.Validate(s)
}
