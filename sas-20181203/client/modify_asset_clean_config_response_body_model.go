// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAssetCleanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyAssetCleanConfigResponseBody
	GetData() *bool
	SetRequestId(v string) *ModifyAssetCleanConfigResponseBody
	GetRequestId() *string
}

type ModifyAssetCleanConfigResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 952776BD-5546-59FC-8AF3-B54EBAD57***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAssetCleanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAssetCleanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAssetCleanConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyAssetCleanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAssetCleanConfigResponseBody) SetData(v bool) *ModifyAssetCleanConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyAssetCleanConfigResponseBody) SetRequestId(v string) *ModifyAssetCleanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAssetCleanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
