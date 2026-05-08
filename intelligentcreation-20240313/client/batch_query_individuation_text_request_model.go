// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryIndividuationTextRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTextIdList(v []*string) *BatchQueryIndividuationTextRequest
	GetTextIdList() []*string
}

type BatchQueryIndividuationTextRequest struct {
	TextIdList []*string `json:"textIdList,omitempty" xml:"textIdList,omitempty" type:"Repeated"`
}

func (s BatchQueryIndividuationTextRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryIndividuationTextRequest) GoString() string {
	return s.String()
}

func (s *BatchQueryIndividuationTextRequest) GetTextIdList() []*string {
	return s.TextIdList
}

func (s *BatchQueryIndividuationTextRequest) SetTextIdList(v []*string) *BatchQueryIndividuationTextRequest {
	s.TextIdList = v
	return s
}

func (s *BatchQueryIndividuationTextRequest) Validate() error {
	return dara.Validate(s)
}
