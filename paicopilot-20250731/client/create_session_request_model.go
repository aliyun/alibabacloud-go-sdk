// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v []*Label) *CreateSessionRequest
	GetLabels() []*Label
	SetTitle(v string) *CreateSessionRequest
	GetTitle() *string
}

type CreateSessionRequest struct {
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Title  *string  `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) GetLabels() []*Label {
	return s.Labels
}

func (s *CreateSessionRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateSessionRequest) SetLabels(v []*Label) *CreateSessionRequest {
	s.Labels = v
	return s
}

func (s *CreateSessionRequest) SetTitle(v string) *CreateSessionRequest {
	s.Title = &v
	return s
}

func (s *CreateSessionRequest) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
