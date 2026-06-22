// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVulWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVulWhitelistResponseBody
	GetRequestId() *string
}

type DeleteVulWhitelistResponseBody struct {
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 9FBC6E47-7508-58C9-9E76-528E118CB1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVulWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVulWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVulWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVulWhitelistResponseBody) SetRequestId(v string) *DeleteVulWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVulWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
