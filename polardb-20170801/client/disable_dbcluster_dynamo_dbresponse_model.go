// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterDynamoDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableDBClusterDynamoDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableDBClusterDynamoDBResponse
	GetStatusCode() *int32
	SetBody(v *DisableDBClusterDynamoDBResponseBody) *DisableDBClusterDynamoDBResponse
	GetBody() *DisableDBClusterDynamoDBResponseBody
}

type DisableDBClusterDynamoDBResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableDBClusterDynamoDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableDBClusterDynamoDBResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterDynamoDBResponse) GoString() string {
	return s.String()
}

func (s *DisableDBClusterDynamoDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableDBClusterDynamoDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableDBClusterDynamoDBResponse) GetBody() *DisableDBClusterDynamoDBResponseBody {
	return s.Body
}

func (s *DisableDBClusterDynamoDBResponse) SetHeaders(v map[string]*string) *DisableDBClusterDynamoDBResponse {
	s.Headers = v
	return s
}

func (s *DisableDBClusterDynamoDBResponse) SetStatusCode(v int32) *DisableDBClusterDynamoDBResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableDBClusterDynamoDBResponse) SetBody(v *DisableDBClusterDynamoDBResponseBody) *DisableDBClusterDynamoDBResponse {
	s.Body = v
	return s
}

func (s *DisableDBClusterDynamoDBResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
