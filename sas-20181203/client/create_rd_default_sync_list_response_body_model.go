// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRdDefaultSyncListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRdDefaultSyncListResponseBody
	GetRequestId() *string
}

type CreateRdDefaultSyncListResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BA674E4B-00CF-5DEA-8B92-360862FB5133
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRdDefaultSyncListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRdDefaultSyncListResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRdDefaultSyncListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRdDefaultSyncListResponseBody) SetRequestId(v string) *CreateRdDefaultSyncListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRdDefaultSyncListResponseBody) Validate() error {
	return dara.Validate(s)
}
