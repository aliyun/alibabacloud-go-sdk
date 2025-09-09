// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupRecycleBinStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetupRecycleBinStatusResponseBody
	GetData() *bool
	SetRequestId(v string) *SetupRecycleBinStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetupRecycleBinStatusResponseBody
	GetSuccess() *bool
}

type SetupRecycleBinStatusResponseBody struct {
	// Indicates whether the table recycle bin is enabled.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A3140FC7-B78B-4D8E-B0C8-926D28******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetupRecycleBinStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetupRecycleBinStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetupRecycleBinStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetupRecycleBinStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetupRecycleBinStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetupRecycleBinStatusResponseBody) SetData(v bool) *SetupRecycleBinStatusResponseBody {
	s.Data = &v
	return s
}

func (s *SetupRecycleBinStatusResponseBody) SetRequestId(v string) *SetupRecycleBinStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupRecycleBinStatusResponseBody) SetSuccess(v bool) *SetupRecycleBinStatusResponseBody {
	s.Success = &v
	return s
}

func (s *SetupRecycleBinStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
