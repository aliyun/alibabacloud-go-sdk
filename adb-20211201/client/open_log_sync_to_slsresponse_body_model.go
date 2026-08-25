// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLogSyncToSLSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenLogSyncToSLSResponseBody
	GetRequestId() *string
}

type OpenLogSyncToSLSResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenLogSyncToSLSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenLogSyncToSLSResponseBody) GoString() string {
	return s.String()
}

func (s *OpenLogSyncToSLSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenLogSyncToSLSResponseBody) SetRequestId(v string) *OpenLogSyncToSLSResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenLogSyncToSLSResponseBody) Validate() error {
	return dara.Validate(s)
}
