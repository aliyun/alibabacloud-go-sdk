// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkAppWebUiAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSparkAppWebUiAddressResponseBodyData) *GetSparkAppWebUiAddressResponseBody
	GetData() *GetSparkAppWebUiAddressResponseBodyData
	SetRequestId(v string) *GetSparkAppWebUiAddressResponseBody
	GetRequestId() *string
}

type GetSparkAppWebUiAddressResponseBody struct {
	// The returned data.
	Data *GetSparkAppWebUiAddressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D65A809F-34CE-4550-9BC1-0ED21ETG380
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkAppWebUiAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppWebUiAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkAppWebUiAddressResponseBody) GetData() *GetSparkAppWebUiAddressResponseBodyData {
	return s.Data
}

func (s *GetSparkAppWebUiAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkAppWebUiAddressResponseBody) SetData(v *GetSparkAppWebUiAddressResponseBodyData) *GetSparkAppWebUiAddressResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBody) SetRequestId(v string) *GetSparkAppWebUiAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSparkAppWebUiAddressResponseBodyData struct {
	// The Spark application ID.
	//
	// example:
	//
	// s202205201533hz1209892000****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The database ID.
	//
	// example:
	//
	// amv-clusterxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The expiration time. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1655801973000
	ExpirationTimeInMillis *int64 `json:"ExpirationTimeInMillis,omitempty" xml:"ExpirationTimeInMillis,omitempty"`
	// The URL of the web UI for the Spark application.
	//
	// example:
	//
	// https://adbsparkui-cn-hangzhou.aliyuncs.com/?token=****
	WebUiAddress *string `json:"WebUiAddress,omitempty" xml:"WebUiAddress,omitempty"`
}

func (s GetSparkAppWebUiAddressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSparkAppWebUiAddressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSparkAppWebUiAddressResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetSparkAppWebUiAddressResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetSparkAppWebUiAddressResponseBodyData) GetExpirationTimeInMillis() *int64 {
	return s.ExpirationTimeInMillis
}

func (s *GetSparkAppWebUiAddressResponseBodyData) GetWebUiAddress() *string {
	return s.WebUiAddress
}

func (s *GetSparkAppWebUiAddressResponseBodyData) SetAppId(v string) *GetSparkAppWebUiAddressResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBodyData) SetDBClusterId(v string) *GetSparkAppWebUiAddressResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBodyData) SetExpirationTimeInMillis(v int64) *GetSparkAppWebUiAddressResponseBodyData {
	s.ExpirationTimeInMillis = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBodyData) SetWebUiAddress(v string) *GetSparkAppWebUiAddressResponseBodyData {
	s.WebUiAddress = &v
	return s
}

func (s *GetSparkAppWebUiAddressResponseBodyData) Validate() error {
	return dara.Validate(s)
}
