// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineRecallManagementServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OfflineRecallManagementServiceResponseBody
	GetRequestId() *string
}

type OfflineRecallManagementServiceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 728C5E01-ABF6-5AA8-B9FC-B3BA05DECC77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OfflineRecallManagementServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineRecallManagementServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineRecallManagementServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineRecallManagementServiceResponseBody) SetRequestId(v string) *OfflineRecallManagementServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineRecallManagementServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
