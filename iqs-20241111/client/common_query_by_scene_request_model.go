// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommonQueryBySceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CommonAgentQuery) *CommonQueryBySceneRequest
	GetBody() *CommonAgentQuery
}

type CommonQueryBySceneRequest struct {
	Body *CommonAgentQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CommonQueryBySceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CommonQueryBySceneRequest) GoString() string {
	return s.String()
}

func (s *CommonQueryBySceneRequest) GetBody() *CommonAgentQuery {
	return s.Body
}

func (s *CommonQueryBySceneRequest) SetBody(v *CommonAgentQuery) *CommonQueryBySceneRequest {
	s.Body = v
	return s
}

func (s *CommonQueryBySceneRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
