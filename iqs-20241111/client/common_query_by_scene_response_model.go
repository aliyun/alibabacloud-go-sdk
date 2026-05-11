// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonQueryBySceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CommonQueryBySceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CommonQueryBySceneResponse
	GetStatusCode() *int32
	SetBody(v *QueryResult) *CommonQueryBySceneResponse
	GetBody() *QueryResult
}

type CommonQueryBySceneResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryResult       `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommonQueryBySceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CommonQueryBySceneResponse) GoString() string {
	return s.String()
}

func (s *CommonQueryBySceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CommonQueryBySceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CommonQueryBySceneResponse) GetBody() *QueryResult {
	return s.Body
}

func (s *CommonQueryBySceneResponse) SetHeaders(v map[string]*string) *CommonQueryBySceneResponse {
	s.Headers = v
	return s
}

func (s *CommonQueryBySceneResponse) SetStatusCode(v int32) *CommonQueryBySceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CommonQueryBySceneResponse) SetBody(v *QueryResult) *CommonQueryBySceneResponse {
	s.Body = v
	return s
}

func (s *CommonQueryBySceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
