// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDashboardNl2sqlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDashboardNl2sqlResponseBody
	GetRequestId() *string
	SetResult(v []*QueryDashboardNl2sqlResponseBodyResult) *QueryDashboardNl2sqlResponseBody
	GetResult() []*QueryDashboardNl2sqlResponseBodyResult
	SetSuccess(v bool) *QueryDashboardNl2sqlResponseBody
	GetSuccess() *bool
}

type QueryDashboardNl2sqlResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 46e537a5****,3dadsu****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution result of the API operation. Valid values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	Result []*QueryDashboardNl2sqlResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Specifies whether the request is successful. Valid values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryDashboardNl2sqlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDashboardNl2sqlResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDashboardNl2sqlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDashboardNl2sqlResponseBody) GetResult() []*QueryDashboardNl2sqlResponseBodyResult {
	return s.Result
}

func (s *QueryDashboardNl2sqlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryDashboardNl2sqlResponseBody) SetRequestId(v string) *QueryDashboardNl2sqlResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDashboardNl2sqlResponseBody) SetResult(v []*QueryDashboardNl2sqlResponseBodyResult) *QueryDashboardNl2sqlResponseBody {
	s.Result = v
	return s
}

func (s *QueryDashboardNl2sqlResponseBody) SetSuccess(v bool) *QueryDashboardNl2sqlResponseBody {
	s.Success = &v
	return s
}

func (s *QueryDashboardNl2sqlResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDashboardNl2sqlResponseBodyResult struct {
	// The permissions on the dashboard NL2SQL resource. If this parameter has a value and contains READ, it indicates that the user has NL2SQL permissions on the dashboard NL2SQL resource.
	Authorities []*string `json:"Authorities,omitempty" xml:"Authorities,omitempty" type:"Repeated"`
	// The name of the dashboard.
	//
	// example:
	//
	// 612b
	DashboardName *string `json:"DashboardName,omitempty" xml:"DashboardName,omitempty"`
	// The NL2SQL resource ID of the dashboard.
	//
	// example:
	//
	// sasdas****sawdau
	DashboardNl2sqlId *string `json:"DashboardNl2sqlId,omitempty" xml:"DashboardNl2sqlId,omitempty"`
	// The user ID of the dashboard creator.
	//
	// example:
	//
	// 46e5374665ba4b679ee22e2a2927****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s QueryDashboardNl2sqlResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s QueryDashboardNl2sqlResponseBodyResult) GoString() string {
	return s.String()
}

func (s *QueryDashboardNl2sqlResponseBodyResult) GetAuthorities() []*string {
	return s.Authorities
}

func (s *QueryDashboardNl2sqlResponseBodyResult) GetDashboardName() *string {
	return s.DashboardName
}

func (s *QueryDashboardNl2sqlResponseBodyResult) GetDashboardNl2sqlId() *string {
	return s.DashboardNl2sqlId
}

func (s *QueryDashboardNl2sqlResponseBodyResult) GetOwnerId() *string {
	return s.OwnerId
}

func (s *QueryDashboardNl2sqlResponseBodyResult) SetAuthorities(v []*string) *QueryDashboardNl2sqlResponseBodyResult {
	s.Authorities = v
	return s
}

func (s *QueryDashboardNl2sqlResponseBodyResult) SetDashboardName(v string) *QueryDashboardNl2sqlResponseBodyResult {
	s.DashboardName = &v
	return s
}

func (s *QueryDashboardNl2sqlResponseBodyResult) SetDashboardNl2sqlId(v string) *QueryDashboardNl2sqlResponseBodyResult {
	s.DashboardNl2sqlId = &v
	return s
}

func (s *QueryDashboardNl2sqlResponseBodyResult) SetOwnerId(v string) *QueryDashboardNl2sqlResponseBodyResult {
	s.OwnerId = &v
	return s
}

func (s *QueryDashboardNl2sqlResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
