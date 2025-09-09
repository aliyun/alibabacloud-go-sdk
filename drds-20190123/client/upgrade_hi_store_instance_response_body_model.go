// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeHiStoreInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpgradeHiStoreInstanceResponseBody
	GetData() *string
	SetRequestId(v string) *UpgradeHiStoreInstanceResponseBody
	GetRequestId() *string
}

type UpgradeHiStoreInstanceResponseBody struct {
	// Indicates whether the request was successful. A value of true indicates that the request was successful. An error message was returned if the request failed.
	//
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DSSDF-SEWE-23ERW
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeHiStoreInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeHiStoreInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *UpgradeHiStoreInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeHiStoreInstanceResponseBody) SetData(v string) *UpgradeHiStoreInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpgradeHiStoreInstanceResponseBody) SetRequestId(v string) *UpgradeHiStoreInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeHiStoreInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
