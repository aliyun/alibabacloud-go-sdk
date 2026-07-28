// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeApplicationVersionResponseBody
	GetRequestId() *string
}

type UpgradeApplicationVersionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeApplicationVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeApplicationVersionResponseBody) SetRequestId(v string) *UpgradeApplicationVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeApplicationVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
