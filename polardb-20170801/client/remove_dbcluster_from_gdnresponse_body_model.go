// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveDBClusterFromGDNResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveDBClusterFromGDNResponseBody
	GetRequestId() *string
}

type RemoveDBClusterFromGDNResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 67F2E75F-AE67-4FB2-821F-A81237EACD15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveDBClusterFromGDNResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveDBClusterFromGDNResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDBClusterFromGDNResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveDBClusterFromGDNResponseBody) SetRequestId(v string) *RemoveDBClusterFromGDNResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDBClusterFromGDNResponseBody) Validate() error {
	return dara.Validate(s)
}
