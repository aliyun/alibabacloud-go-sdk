// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudGroupExpansionServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHybridCloudGroupExpansionServerResponseBody
	GetRequestId() *string
}

type ModifyHybridCloudGroupExpansionServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BD7F9A1-0E72-5BC1-8248-C629B4***B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyHybridCloudGroupExpansionServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudGroupExpansionServerResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudGroupExpansionServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHybridCloudGroupExpansionServerResponseBody) SetRequestId(v string) *ModifyHybridCloudGroupExpansionServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHybridCloudGroupExpansionServerResponseBody) Validate() error {
	return dara.Validate(s)
}
