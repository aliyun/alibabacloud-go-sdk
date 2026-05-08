// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndividuationTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTextIdList(v []*string) *DeleteIndividuationTextRequest
	GetTextIdList() []*string
}

type DeleteIndividuationTextRequest struct {
	TextIdList []*string `json:"textIdList,omitempty" xml:"textIdList,omitempty" type:"Repeated"`
}

func (s DeleteIndividuationTextRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndividuationTextRequest) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationTextRequest) GetTextIdList() []*string {
	return s.TextIdList
}

func (s *DeleteIndividuationTextRequest) SetTextIdList(v []*string) *DeleteIndividuationTextRequest {
	s.TextIdList = v
	return s
}

func (s *DeleteIndividuationTextRequest) Validate() error {
	return dara.Validate(s)
}
