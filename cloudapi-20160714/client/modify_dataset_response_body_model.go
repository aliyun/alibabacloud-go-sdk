// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDatasetResponseBody
	GetRequestId() *string
}

type ModifyDatasetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDatasetResponseBody) SetRequestId(v string) *ModifyDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
