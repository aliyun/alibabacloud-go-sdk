// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpgradeInstanceVersionResponseBody
	GetData() *string
	SetRequestId(v string) *UpgradeInstanceVersionResponseBody
	GetRequestId() *string
}

type UpgradeInstanceVersionResponseBody struct {
	// The result of the request.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2F7F8080-9132-4279-85D0-B7E5C4305162
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeInstanceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *UpgradeInstanceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeInstanceVersionResponseBody) SetData(v string) *UpgradeInstanceVersionResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) SetRequestId(v string) *UpgradeInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeInstanceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
