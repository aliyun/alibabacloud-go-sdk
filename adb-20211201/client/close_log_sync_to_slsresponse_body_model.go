// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseLogSyncToSLSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseLogSyncToSLSResponseBody
	GetRequestId() *string
}

type CloseLogSyncToSLSResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseLogSyncToSLSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseLogSyncToSLSResponseBody) GoString() string {
	return s.String()
}

func (s *CloseLogSyncToSLSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseLogSyncToSLSResponseBody) SetRequestId(v string) *CloseLogSyncToSLSResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseLogSyncToSLSResponseBody) Validate() error {
	return dara.Validate(s)
}
