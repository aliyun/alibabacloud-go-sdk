// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisDimensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody
	GetDatabases() []*string
	SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody
	GetRequestId() *string
	SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody
	GetUserNames() []*string
}

type DescribeDiagnosisDimensionsResponseBody struct {
	// The names of the databases.
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9ADCAACA-E0E8-5319-AE3B-E260E957BDF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The names of the database accounts.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisDimensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponseBody) GetDatabases() []*string {
	return s.Databases
}

func (s *DescribeDiagnosisDimensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisDimensionsResponseBody) GetUserNames() []*string {
	return s.UserNames
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.UserNames = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
