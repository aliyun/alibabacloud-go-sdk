// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustPtsSceneSpeedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AdjustPtsSceneSpeedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AdjustPtsSceneSpeedResponse
	GetStatusCode() *int32
	SetBody(v *AdjustPtsSceneSpeedResponseBody) *AdjustPtsSceneSpeedResponse
	GetBody() *AdjustPtsSceneSpeedResponseBody
}

type AdjustPtsSceneSpeedResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AdjustPtsSceneSpeedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AdjustPtsSceneSpeedResponse) String() string {
	return dara.Prettify(s)
}

func (s AdjustPtsSceneSpeedResponse) GoString() string {
	return s.String()
}

func (s *AdjustPtsSceneSpeedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AdjustPtsSceneSpeedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AdjustPtsSceneSpeedResponse) GetBody() *AdjustPtsSceneSpeedResponseBody {
	return s.Body
}

func (s *AdjustPtsSceneSpeedResponse) SetHeaders(v map[string]*string) *AdjustPtsSceneSpeedResponse {
	s.Headers = v
	return s
}

func (s *AdjustPtsSceneSpeedResponse) SetStatusCode(v int32) *AdjustPtsSceneSpeedResponse {
	s.StatusCode = &v
	return s
}

func (s *AdjustPtsSceneSpeedResponse) SetBody(v *AdjustPtsSceneSpeedResponseBody) *AdjustPtsSceneSpeedResponse {
	s.Body = v
	return s
}

func (s *AdjustPtsSceneSpeedResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
