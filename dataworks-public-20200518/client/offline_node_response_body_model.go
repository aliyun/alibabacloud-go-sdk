// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OfflineNodeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *OfflineNodeResponseBody
	GetSuccess() *string
}

type OfflineNodeResponseBody struct {
	// The request ID. You can use the request ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OfflineNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineNodeResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineNodeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *OfflineNodeResponseBody) SetRequestId(v string) *OfflineNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineNodeResponseBody) SetSuccess(v string) *OfflineNodeResponseBody {
	s.Success = &v
	return s
}

func (s *OfflineNodeResponseBody) Validate() error {
	return dara.Validate(s)
}
