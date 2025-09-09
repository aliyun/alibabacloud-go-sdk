// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRecycleBinTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RemoveRecycleBinTableResponseBody
	GetData() *bool
	SetRequestId(v string) *RemoveRecycleBinTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveRecycleBinTableResponseBody
	GetSuccess() *bool
}

type RemoveRecycleBinTableResponseBody struct {
	// Indicates whether the table in the recycle bin is deleted.
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

func (s RemoveRecycleBinTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveRecycleBinTableResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableResponseBody) GetData() *bool {
	return s.Data
}

func (s *RemoveRecycleBinTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveRecycleBinTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveRecycleBinTableResponseBody) SetData(v bool) *RemoveRecycleBinTableResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveRecycleBinTableResponseBody) SetRequestId(v string) *RemoveRecycleBinTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveRecycleBinTableResponseBody) SetSuccess(v bool) *RemoveRecycleBinTableResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveRecycleBinTableResponseBody) Validate() error {
	return dara.Validate(s)
}
