// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBackupAutoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OpenBackupAutoConfigResponseBody
	GetRequestId() *string
}

type OpenBackupAutoConfigResponseBody struct {
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// ADE57832-9666-511C-9A80-B87DE2E8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenBackupAutoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenBackupAutoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *OpenBackupAutoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenBackupAutoConfigResponseBody) SetRequestId(v string) *OpenBackupAutoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenBackupAutoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
