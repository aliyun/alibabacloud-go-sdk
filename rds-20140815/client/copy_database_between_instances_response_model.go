// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDatabaseBetweenInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CopyDatabaseBetweenInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CopyDatabaseBetweenInstancesResponse
	GetStatusCode() *int32
	SetBody(v *CopyDatabaseBetweenInstancesResponseBody) *CopyDatabaseBetweenInstancesResponse
	GetBody() *CopyDatabaseBetweenInstancesResponseBody
}

type CopyDatabaseBetweenInstancesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyDatabaseBetweenInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyDatabaseBetweenInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s CopyDatabaseBetweenInstancesResponse) GoString() string {
	return s.String()
}

func (s *CopyDatabaseBetweenInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CopyDatabaseBetweenInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CopyDatabaseBetweenInstancesResponse) GetBody() *CopyDatabaseBetweenInstancesResponseBody {
	return s.Body
}

func (s *CopyDatabaseBetweenInstancesResponse) SetHeaders(v map[string]*string) *CopyDatabaseBetweenInstancesResponse {
	s.Headers = v
	return s
}

func (s *CopyDatabaseBetweenInstancesResponse) SetStatusCode(v int32) *CopyDatabaseBetweenInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesResponse) SetBody(v *CopyDatabaseBetweenInstancesResponseBody) *CopyDatabaseBetweenInstancesResponse {
	s.Body = v
	return s
}

func (s *CopyDatabaseBetweenInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
