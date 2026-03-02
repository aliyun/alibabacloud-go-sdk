// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenServiceGroupForServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *OpenServiceGroupForServiceCmd) *OpenServiceGroupForServiceRequest
	GetBody() *OpenServiceGroupForServiceCmd
}

type OpenServiceGroupForServiceRequest struct {
	Body *OpenServiceGroupForServiceCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenServiceGroupForServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenServiceGroupForServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenServiceGroupForServiceRequest) GetBody() *OpenServiceGroupForServiceCmd {
	return s.Body
}

func (s *OpenServiceGroupForServiceRequest) SetBody(v *OpenServiceGroupForServiceCmd) *OpenServiceGroupForServiceRequest {
	s.Body = v
	return s
}

func (s *OpenServiceGroupForServiceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
