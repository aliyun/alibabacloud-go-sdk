// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v int64) *RemoveDataSetRequest
	GetBody() *int64
}

type RemoveDataSetRequest struct {
	Body *int64 `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveDataSetRequest) GoString() string {
	return s.String()
}

func (s *RemoveDataSetRequest) GetBody() *int64 {
	return s.Body
}

func (s *RemoveDataSetRequest) SetBody(v int64) *RemoveDataSetRequest {
	s.Body = &v
	return s
}

func (s *RemoveDataSetRequest) Validate() error {
	return dara.Validate(s)
}
