// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCoordinationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCoIds(v []*string) *CancelCoordinationRequest
	GetCoIds() []*string
}

type CancelCoordinationRequest struct {
	CoIds []*string `json:"CoIds,omitempty" xml:"CoIds,omitempty" type:"Repeated"`
}

func (s CancelCoordinationRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelCoordinationRequest) GoString() string {
	return s.String()
}

func (s *CancelCoordinationRequest) GetCoIds() []*string {
	return s.CoIds
}

func (s *CancelCoordinationRequest) SetCoIds(v []*string) *CancelCoordinationRequest {
	s.CoIds = v
	return s
}

func (s *CancelCoordinationRequest) Validate() error {
	return dara.Validate(s)
}
