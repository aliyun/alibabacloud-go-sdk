// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerScanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyContainerScanConfigResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *ModifyContainerScanConfigResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModifyContainerScanConfigResponseBody
	GetRequestId() *string
}

type ModifyContainerScanConfigResponseBody struct {
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BFF355BC-8A40-55F3-8CBC-CC3E9DAC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyContainerScanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerScanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContainerScanConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyContainerScanConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyContainerScanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyContainerScanConfigResponseBody) SetData(v bool) *ModifyContainerScanConfigResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyContainerScanConfigResponseBody) SetHttpStatusCode(v int32) *ModifyContainerScanConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyContainerScanConfigResponseBody) SetRequestId(v string) *ModifyContainerScanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyContainerScanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
