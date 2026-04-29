// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUnknownThreatDetectProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProcessIdList(v []*string) *DeleteUnknownThreatDetectProcessRequest
	GetProcessIdList() []*string
}

type DeleteUnknownThreatDetectProcessRequest struct {
	ProcessIdList []*string `json:"ProcessIdList,omitempty" xml:"ProcessIdList,omitempty" type:"Repeated"`
}

func (s DeleteUnknownThreatDetectProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUnknownThreatDetectProcessRequest) GoString() string {
	return s.String()
}

func (s *DeleteUnknownThreatDetectProcessRequest) GetProcessIdList() []*string {
	return s.ProcessIdList
}

func (s *DeleteUnknownThreatDetectProcessRequest) SetProcessIdList(v []*string) *DeleteUnknownThreatDetectProcessRequest {
	s.ProcessIdList = v
	return s
}

func (s *DeleteUnknownThreatDetectProcessRequest) Validate() error {
	return dara.Validate(s)
}
