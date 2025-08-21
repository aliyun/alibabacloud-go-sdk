// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatusCodeCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainStatusCodeCountResponseBody
	GetRequestId() *string
	SetStatus200(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus200() *int64
	SetStatus2XX(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus2XX() *int64
	SetStatus3XX(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus3XX() *int64
	SetStatus403(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus403() *int64
	SetStatus404(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus404() *int64
	SetStatus405(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus405() *int64
	SetStatus410(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus410() *int64
	SetStatus499(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus499() *int64
	SetStatus4XX(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus4XX() *int64
	SetStatus501(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus501() *int64
	SetStatus502(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus502() *int64
	SetStatus503(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus503() *int64
	SetStatus504(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus504() *int64
	SetStatus5XX(v int64) *DescribeDomainStatusCodeCountResponseBody
	GetStatus5XX() *int64
}

type DescribeDomainStatusCodeCountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of 200 status codes within the specified period of time.
	//
	// example:
	//
	// 951159
	Status200 *int64 `json:"Status200,omitempty" xml:"Status200,omitempty"`
	// The number of 2xx status codes within the specified period of time.
	//
	// example:
	//
	// 951472
	Status2XX *int64 `json:"Status2XX,omitempty" xml:"Status2XX,omitempty"`
	// The number of 3xx status codes within the specified period of time.
	//
	// example:
	//
	// 133209
	Status3XX *int64 `json:"Status3XX,omitempty" xml:"Status3XX,omitempty"`
	// The number of 403 status codes within the specified period of time.
	//
	// example:
	//
	// 0
	Status403 *int64 `json:"Status403,omitempty" xml:"Status403,omitempty"`
	// The number of 404 status codes within the specified period of time.
	//
	// example:
	//
	// 897
	Status404 *int64 `json:"Status404,omitempty" xml:"Status404,omitempty"`
	// The number of 405 status codes within the specified period of time.
	//
	// example:
	//
	// 0
	Status405 *int64 `json:"Status405,omitempty" xml:"Status405,omitempty"`
	Status410 *int64 `json:"Status410,omitempty" xml:"Status410,omitempty"`
	Status499 *int64 `json:"Status499,omitempty" xml:"Status499,omitempty"`
	// The number of 4xx status codes within the specified period of time.
	//
	// example:
	//
	// 5653
	Status4XX *int64 `json:"Status4XX,omitempty" xml:"Status4XX,omitempty"`
	// The number of 501 status codes within the specified period of time.
	//
	// example:
	//
	// 0
	Status501 *int64 `json:"Status501,omitempty" xml:"Status501,omitempty"`
	// The number of 502 status codes within the specified period of time.
	//
	// example:
	//
	// 0
	Status502 *int64 `json:"Status502,omitempty" xml:"Status502,omitempty"`
	// The number of 503 status codes within the specified period of time.
	//
	// example:
	//
	// 0
	Status503 *int64 `json:"Status503,omitempty" xml:"Status503,omitempty"`
	// The number of 504 status codes within the specified period of time.
	//
	// example:
	//
	// 0
	Status504 *int64 `json:"Status504,omitempty" xml:"Status504,omitempty"`
	// The number of 5xx status codes within the specified period of time.
	//
	// example:
	//
	// 14
	Status5XX *int64 `json:"Status5XX,omitempty" xml:"Status5XX,omitempty"`
}

func (s DescribeDomainStatusCodeCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatusCodeCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus200() *int64 {
	return s.Status200
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus2XX() *int64 {
	return s.Status2XX
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus3XX() *int64 {
	return s.Status3XX
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus403() *int64 {
	return s.Status403
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus404() *int64 {
	return s.Status404
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus405() *int64 {
	return s.Status405
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus410() *int64 {
	return s.Status410
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus499() *int64 {
	return s.Status499
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus4XX() *int64 {
	return s.Status4XX
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus501() *int64 {
	return s.Status501
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus502() *int64 {
	return s.Status502
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus503() *int64 {
	return s.Status503
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus504() *int64 {
	return s.Status504
}

func (s *DescribeDomainStatusCodeCountResponseBody) GetStatus5XX() *int64 {
	return s.Status5XX
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetRequestId(v string) *DescribeDomainStatusCodeCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus200(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status200 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus2XX(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status2XX = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus3XX(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status3XX = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus403(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status403 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus404(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status404 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus405(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status405 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus410(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status410 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus499(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status499 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus4XX(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status4XX = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus501(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status501 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus502(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status502 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus503(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status503 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus504(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status504 = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) SetStatus5XX(v int64) *DescribeDomainStatusCodeCountResponseBody {
	s.Status5XX = &v
	return s
}

func (s *DescribeDomainStatusCodeCountResponseBody) Validate() error {
	return dara.Validate(s)
}
