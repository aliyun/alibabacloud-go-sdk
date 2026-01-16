// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnInstallClusterAddonsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UnInstallClusterAddonsResponseBody
	GetClusterId() *string
	SetRequestId(v string) *UnInstallClusterAddonsResponseBody
	GetRequestId() *string
}

type UnInstallClusterAddonsResponseBody struct {
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnInstallClusterAddonsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnInstallClusterAddonsResponseBody) GoString() string {
	return s.String()
}

func (s *UnInstallClusterAddonsResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UnInstallClusterAddonsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnInstallClusterAddonsResponseBody) SetClusterId(v string) *UnInstallClusterAddonsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UnInstallClusterAddonsResponseBody) SetRequestId(v string) *UnInstallClusterAddonsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnInstallClusterAddonsResponseBody) Validate() error {
	return dara.Validate(s)
}
