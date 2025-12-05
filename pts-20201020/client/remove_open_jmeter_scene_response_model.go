// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOpenJMeterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveOpenJMeterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveOpenJMeterSceneResponse
	GetStatusCode() *int32
	SetBody(v *RemoveOpenJMeterSceneResponseBody) *RemoveOpenJMeterSceneResponse
	GetBody() *RemoveOpenJMeterSceneResponseBody
}

type RemoveOpenJMeterSceneResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveOpenJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveOpenJMeterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveOpenJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *RemoveOpenJMeterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveOpenJMeterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveOpenJMeterSceneResponse) GetBody() *RemoveOpenJMeterSceneResponseBody {
	return s.Body
}

func (s *RemoveOpenJMeterSceneResponse) SetHeaders(v map[string]*string) *RemoveOpenJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *RemoveOpenJMeterSceneResponse) SetStatusCode(v int32) *RemoveOpenJMeterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponse) SetBody(v *RemoveOpenJMeterSceneResponseBody) *RemoveOpenJMeterSceneResponse {
	s.Body = v
	return s
}

func (s *RemoveOpenJMeterSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
