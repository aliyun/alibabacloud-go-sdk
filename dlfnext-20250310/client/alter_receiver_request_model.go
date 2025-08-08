// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlterReceiverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *AlterReceiverRequest
	GetComment() *string
	SetReceiverName(v string) *AlterReceiverRequest
	GetReceiverName() *string
}

type AlterReceiverRequest struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// receiver_name
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
}

func (s AlterReceiverRequest) String() string {
	return dara.Prettify(s)
}

func (s AlterReceiverRequest) GoString() string {
	return s.String()
}

func (s *AlterReceiverRequest) GetComment() *string {
	return s.Comment
}

func (s *AlterReceiverRequest) GetReceiverName() *string {
	return s.ReceiverName
}

func (s *AlterReceiverRequest) SetComment(v string) *AlterReceiverRequest {
	s.Comment = &v
	return s
}

func (s *AlterReceiverRequest) SetReceiverName(v string) *AlterReceiverRequest {
	s.ReceiverName = &v
	return s
}

func (s *AlterReceiverRequest) Validate() error {
	return dara.Validate(s)
}
