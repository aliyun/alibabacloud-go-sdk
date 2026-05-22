// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEdgeContainerAppRecordResponseBody
	GetRequestId() *string
}

type CreateEdgeContainerAppRecordResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEdgeContainerAppRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEdgeContainerAppRecordResponseBody) SetRequestId(v string) *CreateEdgeContainerAppRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEdgeContainerAppRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
