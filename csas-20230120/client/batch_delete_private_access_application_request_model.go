// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePrivateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIds(v []*string) *BatchDeletePrivateAccessApplicationRequest
	GetApplicationIds() []*string
}

type BatchDeletePrivateAccessApplicationRequest struct {
	// The IDs of internal-facing access applications. You can specify up to 100 application IDs.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
}

func (s BatchDeletePrivateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *BatchDeletePrivateAccessApplicationRequest) GetApplicationIds() []*string {
	return s.ApplicationIds
}

func (s *BatchDeletePrivateAccessApplicationRequest) SetApplicationIds(v []*string) *BatchDeletePrivateAccessApplicationRequest {
	s.ApplicationIds = v
	return s
}

func (s *BatchDeletePrivateAccessApplicationRequest) Validate() error {
	return dara.Validate(s)
}
