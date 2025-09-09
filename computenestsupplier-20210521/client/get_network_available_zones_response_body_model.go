// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkAvailableZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetNetworkAvailableZonesResponseBodyData) *GetNetworkAvailableZonesResponseBody
	GetData() *GetNetworkAvailableZonesResponseBodyData
	SetRequestId(v string) *GetNetworkAvailableZonesResponseBody
	GetRequestId() *string
}

type GetNetworkAvailableZonesResponseBody struct {
	Data *GetNetworkAvailableZonesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E8EF75BC-14E4-597A-BE66-FFA9393C0875
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNetworkAvailableZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAvailableZonesResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkAvailableZonesResponseBody) GetData() *GetNetworkAvailableZonesResponseBodyData {
	return s.Data
}

func (s *GetNetworkAvailableZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkAvailableZonesResponseBody) SetData(v *GetNetworkAvailableZonesResponseBodyData) *GetNetworkAvailableZonesResponseBody {
	s.Data = v
	return s
}

func (s *GetNetworkAvailableZonesResponseBody) SetRequestId(v string) *GetNetworkAvailableZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkAvailableZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetNetworkAvailableZonesResponseBodyData struct {
	ZoneIdList []*string `json:"ZoneIdList,omitempty" xml:"ZoneIdList,omitempty" type:"Repeated"`
}

func (s GetNetworkAvailableZonesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAvailableZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetNetworkAvailableZonesResponseBodyData) GetZoneIdList() []*string {
	return s.ZoneIdList
}

func (s *GetNetworkAvailableZonesResponseBodyData) SetZoneIdList(v []*string) *GetNetworkAvailableZonesResponseBodyData {
	s.ZoneIdList = v
	return s
}

func (s *GetNetworkAvailableZonesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
