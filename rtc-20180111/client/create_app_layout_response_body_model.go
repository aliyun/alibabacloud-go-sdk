// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayoutId(v string) *CreateAppLayoutResponseBody
	GetLayoutId() *string
	SetRequestId(v string) *CreateAppLayoutResponseBody
	GetRequestId() *string
}

type CreateAppLayoutResponseBody struct {
	// example:
	//
	// 167466539798442****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 20A6D1E3-1F5F-5440-A4F1-EC7831646FE4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppLayoutResponseBody) GetLayoutId() *string {
	return s.LayoutId
}

func (s *CreateAppLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppLayoutResponseBody) SetLayoutId(v string) *CreateAppLayoutResponseBody {
	s.LayoutId = &v
	return s
}

func (s *CreateAppLayoutResponseBody) SetRequestId(v string) *CreateAppLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
