// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetContentAnalyzeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetContentAnalyzeConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetContentAnalyzeConfigResponseBody
	GetSuccess() *bool
}

type SetContentAnalyzeConfigResponseBody struct {
	// example:
	//
	// 953CFD27-4A2C-54AD-857F-B79EF3A338E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetContentAnalyzeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetContentAnalyzeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetContentAnalyzeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetContentAnalyzeConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetContentAnalyzeConfigResponseBody) SetRequestId(v string) *SetContentAnalyzeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetContentAnalyzeConfigResponseBody) SetSuccess(v bool) *SetContentAnalyzeConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SetContentAnalyzeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
