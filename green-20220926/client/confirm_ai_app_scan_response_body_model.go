// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmAiAppScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ConfirmAiAppScanResponseBody
	GetData() *bool
	SetRequestId(v string) *ConfirmAiAppScanResponseBody
	GetRequestId() *string
}

type ConfirmAiAppScanResponseBody struct {
	// The response data.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID assigned by the backend to uniquely identify the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfirmAiAppScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfirmAiAppScanResponseBody) GoString() string {
	return s.String()
}

func (s *ConfirmAiAppScanResponseBody) GetData() *bool {
	return s.Data
}

func (s *ConfirmAiAppScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfirmAiAppScanResponseBody) SetData(v bool) *ConfirmAiAppScanResponseBody {
	s.Data = &v
	return s
}

func (s *ConfirmAiAppScanResponseBody) SetRequestId(v string) *ConfirmAiAppScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfirmAiAppScanResponseBody) Validate() error {
	return dara.Validate(s)
}
