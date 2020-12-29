// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeAfsTotalConfDataRequest struct {
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataRequest) SetProductName(v string) *DescribeAfsTotalConfDataRequest {
	s.ProductName = &v
	return s
}

type DescribeAfsTotalConfDataResponse struct {
	RequestId                  *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode                    *string                                                       `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	HasData                    *bool                                                         `json:"HasData,omitempty" xml:"HasData,omitempty" require:"true"`
	NcTotalConfVerifyDatas     []*DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas     `json:"NcTotalConfVerifyDatas,omitempty" xml:"NcTotalConfVerifyDatas,omitempty" require:"true" type:"Repeated"`
	NcTotalConfSigDatas        []*DescribeAfsTotalConfDataResponseNcTotalConfSigDatas        `json:"NcTotalConfSigDatas,omitempty" xml:"NcTotalConfSigDatas,omitempty" require:"true" type:"Repeated"`
	NcTotalConfBlockDatas      []*DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas      `json:"NcTotalConfBlockDatas,omitempty" xml:"NcTotalConfBlockDatas,omitempty" require:"true" type:"Repeated"`
	IcTotalConfVerifyDatas     []*DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas     `json:"IcTotalConfVerifyDatas,omitempty" xml:"IcTotalConfVerifyDatas,omitempty" require:"true" type:"Repeated"`
	IcTotalConfSecVerifyDatas  []*DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas  `json:"IcTotalConfSecVerifyDatas,omitempty" xml:"IcTotalConfSecVerifyDatas,omitempty" require:"true" type:"Repeated"`
	IcTotalConfSigDatas        []*DescribeAfsTotalConfDataResponseIcTotalConfSigDatas        `json:"IcTotalConfSigDatas,omitempty" xml:"IcTotalConfSigDatas,omitempty" require:"true" type:"Repeated"`
	IcTotalConfBlockDatas      []*DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas      `json:"IcTotalConfBlockDatas,omitempty" xml:"IcTotalConfBlockDatas,omitempty" require:"true" type:"Repeated"`
	NvcTotalConfVerifyDatas    []*DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas    `json:"NvcTotalConfVerifyDatas,omitempty" xml:"NvcTotalConfVerifyDatas,omitempty" require:"true" type:"Repeated"`
	NvcTotalConfSecVerifyDatas []*DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas `json:"NvcTotalConfSecVerifyDatas,omitempty" xml:"NvcTotalConfSecVerifyDatas,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAfsTotalConfDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponse) SetRequestId(v string) *DescribeAfsTotalConfDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetBizCode(v string) *DescribeAfsTotalConfDataResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetHasData(v bool) *DescribeAfsTotalConfDataResponse {
	s.HasData = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetNcTotalConfVerifyDatas(v []*DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas) *DescribeAfsTotalConfDataResponse {
	s.NcTotalConfVerifyDatas = v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetNcTotalConfSigDatas(v []*DescribeAfsTotalConfDataResponseNcTotalConfSigDatas) *DescribeAfsTotalConfDataResponse {
	s.NcTotalConfSigDatas = v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetNcTotalConfBlockDatas(v []*DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas) *DescribeAfsTotalConfDataResponse {
	s.NcTotalConfBlockDatas = v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetIcTotalConfVerifyDatas(v []*DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas) *DescribeAfsTotalConfDataResponse {
	s.IcTotalConfVerifyDatas = v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetIcTotalConfSecVerifyDatas(v []*DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas) *DescribeAfsTotalConfDataResponse {
	s.IcTotalConfSecVerifyDatas = v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetIcTotalConfSigDatas(v []*DescribeAfsTotalConfDataResponseIcTotalConfSigDatas) *DescribeAfsTotalConfDataResponse {
	s.IcTotalConfSigDatas = v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetIcTotalConfBlockDatas(v []*DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas) *DescribeAfsTotalConfDataResponse {
	s.IcTotalConfBlockDatas = v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetNvcTotalConfVerifyDatas(v []*DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas) *DescribeAfsTotalConfDataResponse {
	s.NvcTotalConfVerifyDatas = v
	return s
}

func (s *DescribeAfsTotalConfDataResponse) SetNvcTotalConfSecVerifyDatas(v []*DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas) *DescribeAfsTotalConfDataResponse {
	s.NvcTotalConfSecVerifyDatas = v
	return s
}

type DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseNcTotalConfVerifyDatas {
	s.Value = &v
	return s
}

type DescribeAfsTotalConfDataResponseNcTotalConfSigDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseNcTotalConfSigDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseNcTotalConfSigDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfSigDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseNcTotalConfSigDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfSigDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseNcTotalConfSigDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfSigDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseNcTotalConfSigDatas {
	s.Value = &v
	return s
}

type DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseNcTotalConfBlockDatas {
	s.Value = &v
	return s
}

type DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseIcTotalConfVerifyDatas {
	s.Value = &v
	return s
}

type DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseIcTotalConfSecVerifyDatas {
	s.Value = &v
	return s
}

type DescribeAfsTotalConfDataResponseIcTotalConfSigDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseIcTotalConfSigDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseIcTotalConfSigDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfSigDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseIcTotalConfSigDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfSigDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseIcTotalConfSigDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfSigDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseIcTotalConfSigDatas {
	s.Value = &v
	return s
}

type DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseIcTotalConfBlockDatas {
	s.Value = &v
	return s
}

type DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseNvcTotalConfVerifyDatas {
	s.Value = &v
	return s
}

type DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas struct {
	Time     *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Category *string `json:"Category,omitempty" xml:"Category,omitempty" require:"true"`
	Value    *int64  `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas) SetTime(v string) *DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas) SetCategory(v string) *DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas {
	s.Category = &v
	return s
}

func (s *DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas) SetValue(v int64) *DescribeAfsTotalConfDataResponseNvcTotalConfSecVerifyDatas {
	s.Value = &v
	return s
}

type DescribeAfsOneConfDataRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	AppKey      *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
	Scene       *string `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty" require:"true"`
}

func (s DescribeAfsOneConfDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsOneConfDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeAfsOneConfDataRequest) SetSourceIp(v string) *DescribeAfsOneConfDataRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAfsOneConfDataRequest) SetAppKey(v string) *DescribeAfsOneConfDataRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeAfsOneConfDataRequest) SetScene(v string) *DescribeAfsOneConfDataRequest {
	s.Scene = &v
	return s
}

func (s *DescribeAfsOneConfDataRequest) SetProductName(v string) *DescribeAfsOneConfDataRequest {
	s.ProductName = &v
	return s
}

type DescribeAfsOneConfDataResponse struct {
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode         *string                                          `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	HasData         *bool                                            `json:"HasData,omitempty" xml:"HasData,omitempty" require:"true"`
	NcOneConfDatas  []*DescribeAfsOneConfDataResponseNcOneConfDatas  `json:"NcOneConfDatas,omitempty" xml:"NcOneConfDatas,omitempty" require:"true" type:"Repeated"`
	IcOneConfDatas  []*DescribeAfsOneConfDataResponseIcOneConfDatas  `json:"IcOneConfDatas,omitempty" xml:"IcOneConfDatas,omitempty" require:"true" type:"Repeated"`
	NvcOneConfDatas []*DescribeAfsOneConfDataResponseNvcOneConfDatas `json:"NvcOneConfDatas,omitempty" xml:"NvcOneConfDatas,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAfsOneConfDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsOneConfDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeAfsOneConfDataResponse) SetRequestId(v string) *DescribeAfsOneConfDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAfsOneConfDataResponse) SetBizCode(v string) *DescribeAfsOneConfDataResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeAfsOneConfDataResponse) SetHasData(v bool) *DescribeAfsOneConfDataResponse {
	s.HasData = &v
	return s
}

func (s *DescribeAfsOneConfDataResponse) SetNcOneConfDatas(v []*DescribeAfsOneConfDataResponseNcOneConfDatas) *DescribeAfsOneConfDataResponse {
	s.NcOneConfDatas = v
	return s
}

func (s *DescribeAfsOneConfDataResponse) SetIcOneConfDatas(v []*DescribeAfsOneConfDataResponseIcOneConfDatas) *DescribeAfsOneConfDataResponse {
	s.IcOneConfDatas = v
	return s
}

func (s *DescribeAfsOneConfDataResponse) SetNvcOneConfDatas(v []*DescribeAfsOneConfDataResponseNvcOneConfDatas) *DescribeAfsOneConfDataResponse {
	s.NvcOneConfDatas = v
	return s
}

type DescribeAfsOneConfDataResponseNcOneConfDatas struct {
	TableDate        *string `json:"TableDate,omitempty" xml:"TableDate,omitempty" require:"true"`
	NcInitCnt        *int    `json:"NcInitCnt,omitempty" xml:"NcInitCnt,omitempty" require:"true"`
	NcNoActionCnt    *int64  `json:"NcNoActionCnt,omitempty" xml:"NcNoActionCnt,omitempty" require:"true"`
	NcVerifyCnt      *int64  `json:"NcVerifyCnt,omitempty" xml:"NcVerifyCnt,omitempty" require:"true"`
	NcVerifyBlockCnt *int64  `json:"NcVerifyBlockCnt,omitempty" xml:"NcVerifyBlockCnt,omitempty" require:"true"`
	NcSigCnt         *int64  `json:"NcSigCnt,omitempty" xml:"NcSigCnt,omitempty" require:"true"`
	NcSigBlockCnt    *int64  `json:"NcSigBlockCnt,omitempty" xml:"NcSigBlockCnt,omitempty" require:"true"`
}

func (s DescribeAfsOneConfDataResponseNcOneConfDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsOneConfDataResponseNcOneConfDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsOneConfDataResponseNcOneConfDatas) SetTableDate(v string) *DescribeAfsOneConfDataResponseNcOneConfDatas {
	s.TableDate = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNcOneConfDatas) SetNcInitCnt(v int) *DescribeAfsOneConfDataResponseNcOneConfDatas {
	s.NcInitCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNcOneConfDatas) SetNcNoActionCnt(v int64) *DescribeAfsOneConfDataResponseNcOneConfDatas {
	s.NcNoActionCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNcOneConfDatas) SetNcVerifyCnt(v int64) *DescribeAfsOneConfDataResponseNcOneConfDatas {
	s.NcVerifyCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNcOneConfDatas) SetNcVerifyBlockCnt(v int64) *DescribeAfsOneConfDataResponseNcOneConfDatas {
	s.NcVerifyBlockCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNcOneConfDatas) SetNcSigCnt(v int64) *DescribeAfsOneConfDataResponseNcOneConfDatas {
	s.NcSigCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNcOneConfDatas) SetNcSigBlockCnt(v int64) *DescribeAfsOneConfDataResponseNcOneConfDatas {
	s.NcSigBlockCnt = &v
	return s
}

type DescribeAfsOneConfDataResponseIcOneConfDatas struct {
	TableDate      *string `json:"TableDate,omitempty" xml:"TableDate,omitempty" require:"true"`
	IcInitCnt      *int64  `json:"IcInitCnt,omitempty" xml:"IcInitCnt,omitempty" require:"true"`
	IcNoActionCnt  *int64  `json:"IcNoActionCnt,omitempty" xml:"IcNoActionCnt,omitempty" require:"true"`
	IcVerifyCnt    *int64  `json:"IcVerifyCnt,omitempty" xml:"IcVerifyCnt,omitempty" require:"true"`
	IcSecVerifyCnt *int64  `json:"IcSecVerifyCnt,omitempty" xml:"IcSecVerifyCnt,omitempty" require:"true"`
	IcSigCnt       *int64  `json:"IcSigCnt,omitempty" xml:"IcSigCnt,omitempty" require:"true"`
	IcBlockCnt     *int64  `json:"IcBlockCnt,omitempty" xml:"IcBlockCnt,omitempty" require:"true"`
}

func (s DescribeAfsOneConfDataResponseIcOneConfDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsOneConfDataResponseIcOneConfDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsOneConfDataResponseIcOneConfDatas) SetTableDate(v string) *DescribeAfsOneConfDataResponseIcOneConfDatas {
	s.TableDate = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseIcOneConfDatas) SetIcInitCnt(v int64) *DescribeAfsOneConfDataResponseIcOneConfDatas {
	s.IcInitCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseIcOneConfDatas) SetIcNoActionCnt(v int64) *DescribeAfsOneConfDataResponseIcOneConfDatas {
	s.IcNoActionCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseIcOneConfDatas) SetIcVerifyCnt(v int64) *DescribeAfsOneConfDataResponseIcOneConfDatas {
	s.IcVerifyCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseIcOneConfDatas) SetIcSecVerifyCnt(v int64) *DescribeAfsOneConfDataResponseIcOneConfDatas {
	s.IcSecVerifyCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseIcOneConfDatas) SetIcSigCnt(v int64) *DescribeAfsOneConfDataResponseIcOneConfDatas {
	s.IcSigCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseIcOneConfDatas) SetIcBlockCnt(v int64) *DescribeAfsOneConfDataResponseIcOneConfDatas {
	s.IcBlockCnt = &v
	return s
}

type DescribeAfsOneConfDataResponseNvcOneConfDatas struct {
	TableDate       *string `json:"TableDate,omitempty" xml:"TableDate,omitempty" require:"true"`
	NvcInitCnt      *int64  `json:"NvcInitCnt,omitempty" xml:"NvcInitCnt,omitempty" require:"true"`
	NvcNoActionCnt  *int64  `json:"NvcNoActionCnt,omitempty" xml:"NvcNoActionCnt,omitempty" require:"true"`
	NvcVerifyCnt    *int64  `json:"NvcVerifyCnt,omitempty" xml:"NvcVerifyCnt,omitempty" require:"true"`
	NvcSecVerifyCnt *int64  `json:"NvcSecVerifyCnt,omitempty" xml:"NvcSecVerifyCnt,omitempty" require:"true"`
	NvcBlockCnt     *int64  `json:"NvcBlockCnt,omitempty" xml:"NvcBlockCnt,omitempty" require:"true"`
}

func (s DescribeAfsOneConfDataResponseNvcOneConfDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsOneConfDataResponseNvcOneConfDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsOneConfDataResponseNvcOneConfDatas) SetTableDate(v string) *DescribeAfsOneConfDataResponseNvcOneConfDatas {
	s.TableDate = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNvcOneConfDatas) SetNvcInitCnt(v int64) *DescribeAfsOneConfDataResponseNvcOneConfDatas {
	s.NvcInitCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNvcOneConfDatas) SetNvcNoActionCnt(v int64) *DescribeAfsOneConfDataResponseNvcOneConfDatas {
	s.NvcNoActionCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNvcOneConfDatas) SetNvcVerifyCnt(v int64) *DescribeAfsOneConfDataResponseNvcOneConfDatas {
	s.NvcVerifyCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNvcOneConfDatas) SetNvcSecVerifyCnt(v int64) *DescribeAfsOneConfDataResponseNvcOneConfDatas {
	s.NvcSecVerifyCnt = &v
	return s
}

func (s *DescribeAfsOneConfDataResponseNvcOneConfDatas) SetNvcBlockCnt(v int64) *DescribeAfsOneConfDataResponseNvcOneConfDatas {
	s.NvcBlockCnt = &v
	return s
}

type DescribeAfsVerifySigDataRequest struct {
	SourceIp    *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	AppKey      *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
	Scene       *string `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty" require:"true"`
}

func (s DescribeAfsVerifySigDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataRequest) SetSourceIp(v string) *DescribeAfsVerifySigDataRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAfsVerifySigDataRequest) SetAppKey(v string) *DescribeAfsVerifySigDataRequest {
	s.AppKey = &v
	return s
}

func (s *DescribeAfsVerifySigDataRequest) SetScene(v string) *DescribeAfsVerifySigDataRequest {
	s.Scene = &v
	return s
}

func (s *DescribeAfsVerifySigDataRequest) SetProductName(v string) *DescribeAfsVerifySigDataRequest {
	s.ProductName = &v
	return s
}

type DescribeAfsVerifySigDataResponse struct {
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode          *string                                             `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	HasData          *bool                                               `json:"HasData,omitempty" xml:"HasData,omitempty" require:"true"`
	NcVerifyDatas    []*DescribeAfsVerifySigDataResponseNcVerifyDatas    `json:"NcVerifyDatas,omitempty" xml:"NcVerifyDatas,omitempty" require:"true" type:"Repeated"`
	NcSigDatas       []*DescribeAfsVerifySigDataResponseNcSigDatas       `json:"NcSigDatas,omitempty" xml:"NcSigDatas,omitempty" require:"true" type:"Repeated"`
	IcSecVerifyDatas []*DescribeAfsVerifySigDataResponseIcSecVerifyDatas `json:"IcSecVerifyDatas,omitempty" xml:"IcSecVerifyDatas,omitempty" require:"true" type:"Repeated"`
	IcVerifyDatas    []*DescribeAfsVerifySigDataResponseIcVerifyDatas    `json:"IcVerifyDatas,omitempty" xml:"IcVerifyDatas,omitempty" require:"true" type:"Repeated"`
	NvcVerifyDatas   []*DescribeAfsVerifySigDataResponseNvcVerifyDatas   `json:"NvcVerifyDatas,omitempty" xml:"NvcVerifyDatas,omitempty" require:"true" type:"Repeated"`
	NvcSecDatas      []*DescribeAfsVerifySigDataResponseNvcSecDatas      `json:"NvcSecDatas,omitempty" xml:"NvcSecDatas,omitempty" require:"true" type:"Repeated"`
	NvcCodeDatas     []*DescribeAfsVerifySigDataResponseNvcCodeDatas     `json:"NvcCodeDatas,omitempty" xml:"NvcCodeDatas,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAfsVerifySigDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataResponse) SetRequestId(v string) *DescribeAfsVerifySigDataResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetBizCode(v string) *DescribeAfsVerifySigDataResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetHasData(v bool) *DescribeAfsVerifySigDataResponse {
	s.HasData = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetNcVerifyDatas(v []*DescribeAfsVerifySigDataResponseNcVerifyDatas) *DescribeAfsVerifySigDataResponse {
	s.NcVerifyDatas = v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetNcSigDatas(v []*DescribeAfsVerifySigDataResponseNcSigDatas) *DescribeAfsVerifySigDataResponse {
	s.NcSigDatas = v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetIcSecVerifyDatas(v []*DescribeAfsVerifySigDataResponseIcSecVerifyDatas) *DescribeAfsVerifySigDataResponse {
	s.IcSecVerifyDatas = v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetIcVerifyDatas(v []*DescribeAfsVerifySigDataResponseIcVerifyDatas) *DescribeAfsVerifySigDataResponse {
	s.IcVerifyDatas = v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetNvcVerifyDatas(v []*DescribeAfsVerifySigDataResponseNvcVerifyDatas) *DescribeAfsVerifySigDataResponse {
	s.NvcVerifyDatas = v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetNvcSecDatas(v []*DescribeAfsVerifySigDataResponseNvcSecDatas) *DescribeAfsVerifySigDataResponse {
	s.NvcSecDatas = v
	return s
}

func (s *DescribeAfsVerifySigDataResponse) SetNvcCodeDatas(v []*DescribeAfsVerifySigDataResponseNvcCodeDatas) *DescribeAfsVerifySigDataResponse {
	s.NvcCodeDatas = v
	return s
}

type DescribeAfsVerifySigDataResponseNcVerifyDatas struct {
	Time          *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	NcVerifyPass  *int64  `json:"NcVerifyPass,omitempty" xml:"NcVerifyPass,omitempty" require:"true"`
	NcVerifyBlock *int64  `json:"NcVerifyBlock,omitempty" xml:"NcVerifyBlock,omitempty" require:"true"`
}

func (s DescribeAfsVerifySigDataResponseNcVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataResponseNcVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataResponseNcVerifyDatas) SetTime(v string) *DescribeAfsVerifySigDataResponseNcVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNcVerifyDatas) SetNcVerifyPass(v int64) *DescribeAfsVerifySigDataResponseNcVerifyDatas {
	s.NcVerifyPass = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNcVerifyDatas) SetNcVerifyBlock(v int64) *DescribeAfsVerifySigDataResponseNcVerifyDatas {
	s.NcVerifyBlock = &v
	return s
}

type DescribeAfsVerifySigDataResponseNcSigDatas struct {
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	NcSigPass  *int64  `json:"NcSigPass,omitempty" xml:"NcSigPass,omitempty" require:"true"`
	NcSigBlock *int64  `json:"NcSigBlock,omitempty" xml:"NcSigBlock,omitempty" require:"true"`
}

func (s DescribeAfsVerifySigDataResponseNcSigDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataResponseNcSigDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataResponseNcSigDatas) SetTime(v string) *DescribeAfsVerifySigDataResponseNcSigDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNcSigDatas) SetNcSigPass(v int64) *DescribeAfsVerifySigDataResponseNcSigDatas {
	s.NcSigPass = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNcSigDatas) SetNcSigBlock(v int64) *DescribeAfsVerifySigDataResponseNcSigDatas {
	s.NcSigBlock = &v
	return s
}

type DescribeAfsVerifySigDataResponseIcSecVerifyDatas struct {
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	IcSecPass  *int64  `json:"IcSecPass,omitempty" xml:"IcSecPass,omitempty" require:"true"`
	IcSecBlock *int64  `json:"IcSecBlock,omitempty" xml:"IcSecBlock,omitempty" require:"true"`
}

func (s DescribeAfsVerifySigDataResponseIcSecVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataResponseIcSecVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataResponseIcSecVerifyDatas) SetTime(v string) *DescribeAfsVerifySigDataResponseIcSecVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseIcSecVerifyDatas) SetIcSecPass(v int64) *DescribeAfsVerifySigDataResponseIcSecVerifyDatas {
	s.IcSecPass = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseIcSecVerifyDatas) SetIcSecBlock(v int64) *DescribeAfsVerifySigDataResponseIcSecVerifyDatas {
	s.IcSecBlock = &v
	return s
}

type DescribeAfsVerifySigDataResponseIcVerifyDatas struct {
	Time           *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	IcVerifyCnt    *int64  `json:"IcVerifyCnt,omitempty" xml:"IcVerifyCnt,omitempty" require:"true"`
	IcSecVerifyCnt *int64  `json:"IcSecVerifyCnt,omitempty" xml:"IcSecVerifyCnt,omitempty" require:"true"`
	IcSigCnt       *int64  `json:"IcSigCnt,omitempty" xml:"IcSigCnt,omitempty" require:"true"`
	IcBlockCnt     *int64  `json:"IcBlockCnt,omitempty" xml:"IcBlockCnt,omitempty" require:"true"`
}

func (s DescribeAfsVerifySigDataResponseIcVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataResponseIcVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataResponseIcVerifyDatas) SetTime(v string) *DescribeAfsVerifySigDataResponseIcVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseIcVerifyDatas) SetIcVerifyCnt(v int64) *DescribeAfsVerifySigDataResponseIcVerifyDatas {
	s.IcVerifyCnt = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseIcVerifyDatas) SetIcSecVerifyCnt(v int64) *DescribeAfsVerifySigDataResponseIcVerifyDatas {
	s.IcSecVerifyCnt = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseIcVerifyDatas) SetIcSigCnt(v int64) *DescribeAfsVerifySigDataResponseIcVerifyDatas {
	s.IcSigCnt = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseIcVerifyDatas) SetIcBlockCnt(v int64) *DescribeAfsVerifySigDataResponseIcVerifyDatas {
	s.IcBlockCnt = &v
	return s
}

type DescribeAfsVerifySigDataResponseNvcVerifyDatas struct {
	Time            *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	NvcVerifyCnt    *int64  `json:"NvcVerifyCnt,omitempty" xml:"NvcVerifyCnt,omitempty" require:"true"`
	NvcSecVerifyCnt *int64  `json:"NvcSecVerifyCnt,omitempty" xml:"NvcSecVerifyCnt,omitempty" require:"true"`
}

func (s DescribeAfsVerifySigDataResponseNvcVerifyDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataResponseNvcVerifyDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataResponseNvcVerifyDatas) SetTime(v string) *DescribeAfsVerifySigDataResponseNvcVerifyDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNvcVerifyDatas) SetNvcVerifyCnt(v int64) *DescribeAfsVerifySigDataResponseNvcVerifyDatas {
	s.NvcVerifyCnt = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNvcVerifyDatas) SetNvcSecVerifyCnt(v int64) *DescribeAfsVerifySigDataResponseNvcVerifyDatas {
	s.NvcSecVerifyCnt = &v
	return s
}

type DescribeAfsVerifySigDataResponseNvcSecDatas struct {
	Time        *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	NvcSecPass  *int64  `json:"NvcSecPass,omitempty" xml:"NvcSecPass,omitempty" require:"true"`
	NvcSecBlock *int64  `json:"NvcSecBlock,omitempty" xml:"NvcSecBlock,omitempty" require:"true"`
}

func (s DescribeAfsVerifySigDataResponseNvcSecDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataResponseNvcSecDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataResponseNvcSecDatas) SetTime(v string) *DescribeAfsVerifySigDataResponseNvcSecDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNvcSecDatas) SetNvcSecPass(v int64) *DescribeAfsVerifySigDataResponseNvcSecDatas {
	s.NvcSecPass = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNvcSecDatas) SetNvcSecBlock(v int64) *DescribeAfsVerifySigDataResponseNvcSecDatas {
	s.NvcSecBlock = &v
	return s
}

type DescribeAfsVerifySigDataResponseNvcCodeDatas struct {
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	NvcCode200 *int64  `json:"NvcCode200,omitempty" xml:"NvcCode200,omitempty" require:"true"`
	NvcCode400 *int64  `json:"NvcCode400,omitempty" xml:"NvcCode400,omitempty" require:"true"`
	NvcCode600 *int64  `json:"NvcCode600,omitempty" xml:"NvcCode600,omitempty" require:"true"`
	NvcCode800 *int64  `json:"NvcCode800,omitempty" xml:"NvcCode800,omitempty" require:"true"`
}

func (s DescribeAfsVerifySigDataResponseNvcCodeDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsVerifySigDataResponseNvcCodeDatas) GoString() string {
	return s.String()
}

func (s *DescribeAfsVerifySigDataResponseNvcCodeDatas) SetTime(v string) *DescribeAfsVerifySigDataResponseNvcCodeDatas {
	s.Time = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNvcCodeDatas) SetNvcCode200(v int64) *DescribeAfsVerifySigDataResponseNvcCodeDatas {
	s.NvcCode200 = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNvcCodeDatas) SetNvcCode400(v int64) *DescribeAfsVerifySigDataResponseNvcCodeDatas {
	s.NvcCode400 = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNvcCodeDatas) SetNvcCode600(v int64) *DescribeAfsVerifySigDataResponseNvcCodeDatas {
	s.NvcCode600 = &v
	return s
}

func (s *DescribeAfsVerifySigDataResponseNvcCodeDatas) SetNvcCode800(v int64) *DescribeAfsVerifySigDataResponseNvcCodeDatas {
	s.NvcCode800 = &v
	return s
}

type DescribeAfsConfigNameRequest struct {
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty" require:"true"`
}

func (s DescribeAfsConfigNameRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsConfigNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeAfsConfigNameRequest) SetProductName(v string) *DescribeAfsConfigNameRequest {
	s.ProductName = &v
	return s
}

type DescribeAfsConfigNameResponse struct {
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode     *string                                     `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	HasData     *bool                                       `json:"HasData,omitempty" xml:"HasData,omitempty" require:"true"`
	ConfigNames []*DescribeAfsConfigNameResponseConfigNames `json:"ConfigNames,omitempty" xml:"ConfigNames,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAfsConfigNameResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsConfigNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeAfsConfigNameResponse) SetRequestId(v string) *DescribeAfsConfigNameResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAfsConfigNameResponse) SetBizCode(v string) *DescribeAfsConfigNameResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeAfsConfigNameResponse) SetHasData(v bool) *DescribeAfsConfigNameResponse {
	s.HasData = &v
	return s
}

func (s *DescribeAfsConfigNameResponse) SetConfigNames(v []*DescribeAfsConfigNameResponseConfigNames) *DescribeAfsConfigNameResponse {
	s.ConfigNames = v
	return s
}

type DescribeAfsConfigNameResponseConfigNames struct {
	AliUid     *string `json:"AliUid,omitempty" xml:"AliUid,omitempty" require:"true"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty" require:"true"`
	AppKey     *string `json:"AppKey,omitempty" xml:"AppKey,omitempty" require:"true"`
	Scene      *string `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true"`
	RefExtId   *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty" require:"true"`
}

func (s DescribeAfsConfigNameResponseConfigNames) String() string {
	return tea.Prettify(s)
}

func (s DescribeAfsConfigNameResponseConfigNames) GoString() string {
	return s.String()
}

func (s *DescribeAfsConfigNameResponseConfigNames) SetAliUid(v string) *DescribeAfsConfigNameResponseConfigNames {
	s.AliUid = &v
	return s
}

func (s *DescribeAfsConfigNameResponseConfigNames) SetConfigName(v string) *DescribeAfsConfigNameResponseConfigNames {
	s.ConfigName = &v
	return s
}

func (s *DescribeAfsConfigNameResponseConfigNames) SetAppKey(v string) *DescribeAfsConfigNameResponseConfigNames {
	s.AppKey = &v
	return s
}

func (s *DescribeAfsConfigNameResponseConfigNames) SetScene(v string) *DescribeAfsConfigNameResponseConfigNames {
	s.Scene = &v
	return s
}

func (s *DescribeAfsConfigNameResponseConfigNames) SetRefExtId(v string) *DescribeAfsConfigNameResponseConfigNames {
	s.RefExtId = &v
	return s
}

type UpdateConfigNameRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	RefExtId   *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty" require:"true"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty" require:"true"`
}

func (s UpdateConfigNameRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigNameRequest) SetSourceIp(v string) *UpdateConfigNameRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateConfigNameRequest) SetLang(v string) *UpdateConfigNameRequest {
	s.Lang = &v
	return s
}

func (s *UpdateConfigNameRequest) SetRefExtId(v string) *UpdateConfigNameRequest {
	s.RefExtId = &v
	return s
}

func (s *UpdateConfigNameRequest) SetConfigName(v string) *UpdateConfigNameRequest {
	s.ConfigName = &v
	return s
}

type UpdateConfigNameResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode   *string `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
}

func (s UpdateConfigNameResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigNameResponse) SetRequestId(v string) *UpdateConfigNameResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateConfigNameResponse) SetBizCode(v string) *UpdateConfigNameResponse {
	s.BizCode = &v
	return s
}

type DescribeCaptchaOrderRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	Lang     *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeCaptchaOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaOrderRequest) SetSourceIp(v string) *DescribeCaptchaOrderRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCaptchaOrderRequest) SetLang(v string) *DescribeCaptchaOrderRequest {
	s.Lang = &v
	return s
}

type DescribeCaptchaOrderResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode   *string `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
}

func (s DescribeCaptchaOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaOrderResponse) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaOrderResponse) SetRequestId(v string) *DescribeCaptchaOrderResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCaptchaOrderResponse) SetBizCode(v string) *DescribeCaptchaOrderResponse {
	s.BizCode = &v
	return s
}

type DescribeOrderInfoRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeOrderInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrderInfoRequest) SetSourceIp(v string) *DescribeOrderInfoRequest {
	s.SourceIp = &v
	return s
}

type DescribeOrderInfoResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode    *string `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	OrderLevel *string `json:"OrderLevel,omitempty" xml:"OrderLevel,omitempty" require:"true"`
	Num        *string `json:"Num,omitempty" xml:"Num,omitempty" require:"true"`
	BeginDate  *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty" require:"true"`
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty" require:"true"`
}

func (s DescribeOrderInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrderInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrderInfoResponse) SetRequestId(v string) *DescribeOrderInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeOrderInfoResponse) SetBizCode(v string) *DescribeOrderInfoResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeOrderInfoResponse) SetOrderLevel(v string) *DescribeOrderInfoResponse {
	s.OrderLevel = &v
	return s
}

func (s *DescribeOrderInfoResponse) SetNum(v string) *DescribeOrderInfoResponse {
	s.Num = &v
	return s
}

func (s *DescribeOrderInfoResponse) SetBeginDate(v string) *DescribeOrderInfoResponse {
	s.BeginDate = &v
	return s
}

func (s *DescribeOrderInfoResponse) SetEndDate(v string) *DescribeOrderInfoResponse {
	s.EndDate = &v
	return s
}

type SetEarlyWarningRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	WarnOpen  *bool   `json:"WarnOpen,omitempty" xml:"WarnOpen,omitempty" require:"true"`
	Channel   *string `json:"Channel,omitempty" xml:"Channel,omitempty" require:"true"`
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty" require:"true"`
	TimeOpen  *bool   `json:"TimeOpen,omitempty" xml:"TimeOpen,omitempty" require:"true"`
	TimeBegin *string `json:"TimeBegin,omitempty" xml:"TimeBegin,omitempty" require:"true"`
	TimeEnd   *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty" require:"true"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty" require:"true"`
}

func (s SetEarlyWarningRequest) String() string {
	return tea.Prettify(s)
}

func (s SetEarlyWarningRequest) GoString() string {
	return s.String()
}

func (s *SetEarlyWarningRequest) SetSourceIp(v string) *SetEarlyWarningRequest {
	s.SourceIp = &v
	return s
}

func (s *SetEarlyWarningRequest) SetWarnOpen(v bool) *SetEarlyWarningRequest {
	s.WarnOpen = &v
	return s
}

func (s *SetEarlyWarningRequest) SetChannel(v string) *SetEarlyWarningRequest {
	s.Channel = &v
	return s
}

func (s *SetEarlyWarningRequest) SetFrequency(v string) *SetEarlyWarningRequest {
	s.Frequency = &v
	return s
}

func (s *SetEarlyWarningRequest) SetTimeOpen(v bool) *SetEarlyWarningRequest {
	s.TimeOpen = &v
	return s
}

func (s *SetEarlyWarningRequest) SetTimeBegin(v string) *SetEarlyWarningRequest {
	s.TimeBegin = &v
	return s
}

func (s *SetEarlyWarningRequest) SetTimeEnd(v string) *SetEarlyWarningRequest {
	s.TimeEnd = &v
	return s
}

func (s *SetEarlyWarningRequest) SetTitle(v string) *SetEarlyWarningRequest {
	s.Title = &v
	return s
}

type SetEarlyWarningResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode   *string `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
}

func (s SetEarlyWarningResponse) String() string {
	return tea.Prettify(s)
}

func (s SetEarlyWarningResponse) GoString() string {
	return s.String()
}

func (s *SetEarlyWarningResponse) SetRequestId(v string) *SetEarlyWarningResponse {
	s.RequestId = &v
	return s
}

func (s *SetEarlyWarningResponse) SetBizCode(v string) *SetEarlyWarningResponse {
	s.BizCode = &v
	return s
}

type DescribePersonMachineListRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribePersonMachineListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePersonMachineListRequest) GoString() string {
	return s.String()
}

func (s *DescribePersonMachineListRequest) SetSourceIp(v string) *DescribePersonMachineListRequest {
	s.SourceIp = &v
	return s
}

type DescribePersonMachineListResponse struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode          *string                                            `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	PersonMachineRes *DescribePersonMachineListResponsePersonMachineRes `json:"PersonMachineRes,omitempty" xml:"PersonMachineRes,omitempty" require:"true" type:"Struct"`
}

func (s DescribePersonMachineListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePersonMachineListResponse) GoString() string {
	return s.String()
}

func (s *DescribePersonMachineListResponse) SetRequestId(v string) *DescribePersonMachineListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribePersonMachineListResponse) SetBizCode(v string) *DescribePersonMachineListResponse {
	s.BizCode = &v
	return s
}

func (s *DescribePersonMachineListResponse) SetPersonMachineRes(v *DescribePersonMachineListResponsePersonMachineRes) *DescribePersonMachineListResponse {
	s.PersonMachineRes = v
	return s
}

type DescribePersonMachineListResponsePersonMachineRes struct {
	HasConfiguration *string                                                            `json:"HasConfiguration,omitempty" xml:"HasConfiguration,omitempty" require:"true"`
	PersonMachines   []*DescribePersonMachineListResponsePersonMachineResPersonMachines `json:"PersonMachines,omitempty" xml:"PersonMachines,omitempty" require:"true" type:"Repeated"`
}

func (s DescribePersonMachineListResponsePersonMachineRes) String() string {
	return tea.Prettify(s)
}

func (s DescribePersonMachineListResponsePersonMachineRes) GoString() string {
	return s.String()
}

func (s *DescribePersonMachineListResponsePersonMachineRes) SetHasConfiguration(v string) *DescribePersonMachineListResponsePersonMachineRes {
	s.HasConfiguration = &v
	return s
}

func (s *DescribePersonMachineListResponsePersonMachineRes) SetPersonMachines(v []*DescribePersonMachineListResponsePersonMachineResPersonMachines) *DescribePersonMachineListResponsePersonMachineRes {
	s.PersonMachines = v
	return s
}

type DescribePersonMachineListResponsePersonMachineResPersonMachines struct {
	ConfigurationName   *string `json:"ConfigurationName,omitempty" xml:"ConfigurationName,omitempty" require:"true"`
	Appkey              *string `json:"Appkey,omitempty" xml:"Appkey,omitempty" require:"true"`
	ConfigurationMethod *string `json:"ConfigurationMethod,omitempty" xml:"ConfigurationMethod,omitempty" require:"true"`
	ApplyType           *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty" require:"true"`
	Scene               *string `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true"`
	LastUpdate          *string `json:"LastUpdate,omitempty" xml:"LastUpdate,omitempty" require:"true"`
	ExtId               *string `json:"ExtId,omitempty" xml:"ExtId,omitempty" require:"true"`
	SceneOriginal       *string `json:"SceneOriginal,omitempty" xml:"SceneOriginal,omitempty" require:"true"`
}

func (s DescribePersonMachineListResponsePersonMachineResPersonMachines) String() string {
	return tea.Prettify(s)
}

func (s DescribePersonMachineListResponsePersonMachineResPersonMachines) GoString() string {
	return s.String()
}

func (s *DescribePersonMachineListResponsePersonMachineResPersonMachines) SetConfigurationName(v string) *DescribePersonMachineListResponsePersonMachineResPersonMachines {
	s.ConfigurationName = &v
	return s
}

func (s *DescribePersonMachineListResponsePersonMachineResPersonMachines) SetAppkey(v string) *DescribePersonMachineListResponsePersonMachineResPersonMachines {
	s.Appkey = &v
	return s
}

func (s *DescribePersonMachineListResponsePersonMachineResPersonMachines) SetConfigurationMethod(v string) *DescribePersonMachineListResponsePersonMachineResPersonMachines {
	s.ConfigurationMethod = &v
	return s
}

func (s *DescribePersonMachineListResponsePersonMachineResPersonMachines) SetApplyType(v string) *DescribePersonMachineListResponsePersonMachineResPersonMachines {
	s.ApplyType = &v
	return s
}

func (s *DescribePersonMachineListResponsePersonMachineResPersonMachines) SetScene(v string) *DescribePersonMachineListResponsePersonMachineResPersonMachines {
	s.Scene = &v
	return s
}

func (s *DescribePersonMachineListResponsePersonMachineResPersonMachines) SetLastUpdate(v string) *DescribePersonMachineListResponsePersonMachineResPersonMachines {
	s.LastUpdate = &v
	return s
}

func (s *DescribePersonMachineListResponsePersonMachineResPersonMachines) SetExtId(v string) *DescribePersonMachineListResponsePersonMachineResPersonMachines {
	s.ExtId = &v
	return s
}

func (s *DescribePersonMachineListResponsePersonMachineResPersonMachines) SetSceneOriginal(v string) *DescribePersonMachineListResponsePersonMachineResPersonMachines {
	s.SceneOriginal = &v
	return s
}

type DescribeEarlyWarningRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeEarlyWarningRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEarlyWarningRequest) GoString() string {
	return s.String()
}

func (s *DescribeEarlyWarningRequest) SetSourceIp(v string) *DescribeEarlyWarningRequest {
	s.SourceIp = &v
	return s
}

type DescribeEarlyWarningResponse struct {
	RequestId     *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	HasWarning    *bool                                        `json:"HasWarning,omitempty" xml:"HasWarning,omitempty" require:"true"`
	BizCode       *string                                      `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	EarlyWarnings []*DescribeEarlyWarningResponseEarlyWarnings `json:"EarlyWarnings,omitempty" xml:"EarlyWarnings,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeEarlyWarningResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEarlyWarningResponse) GoString() string {
	return s.String()
}

func (s *DescribeEarlyWarningResponse) SetRequestId(v string) *DescribeEarlyWarningResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeEarlyWarningResponse) SetHasWarning(v bool) *DescribeEarlyWarningResponse {
	s.HasWarning = &v
	return s
}

func (s *DescribeEarlyWarningResponse) SetBizCode(v string) *DescribeEarlyWarningResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeEarlyWarningResponse) SetEarlyWarnings(v []*DescribeEarlyWarningResponseEarlyWarnings) *DescribeEarlyWarningResponse {
	s.EarlyWarnings = v
	return s
}

type DescribeEarlyWarningResponseEarlyWarnings struct {
	WarnOpen  *bool   `json:"WarnOpen,omitempty" xml:"WarnOpen,omitempty" require:"true"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty" require:"true"`
	Content   *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true"`
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty" require:"true"`
	TimeOpen  *bool   `json:"TimeOpen,omitempty" xml:"TimeOpen,omitempty" require:"true"`
	TimeBegin *string `json:"TimeBegin,omitempty" xml:"TimeBegin,omitempty" require:"true"`
	TimeEnd   *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty" require:"true"`
	Channel   *string `json:"Channel,omitempty" xml:"Channel,omitempty" require:"true"`
}

func (s DescribeEarlyWarningResponseEarlyWarnings) String() string {
	return tea.Prettify(s)
}

func (s DescribeEarlyWarningResponseEarlyWarnings) GoString() string {
	return s.String()
}

func (s *DescribeEarlyWarningResponseEarlyWarnings) SetWarnOpen(v bool) *DescribeEarlyWarningResponseEarlyWarnings {
	s.WarnOpen = &v
	return s
}

func (s *DescribeEarlyWarningResponseEarlyWarnings) SetTitle(v string) *DescribeEarlyWarningResponseEarlyWarnings {
	s.Title = &v
	return s
}

func (s *DescribeEarlyWarningResponseEarlyWarnings) SetContent(v string) *DescribeEarlyWarningResponseEarlyWarnings {
	s.Content = &v
	return s
}

func (s *DescribeEarlyWarningResponseEarlyWarnings) SetFrequency(v string) *DescribeEarlyWarningResponseEarlyWarnings {
	s.Frequency = &v
	return s
}

func (s *DescribeEarlyWarningResponseEarlyWarnings) SetTimeOpen(v bool) *DescribeEarlyWarningResponseEarlyWarnings {
	s.TimeOpen = &v
	return s
}

func (s *DescribeEarlyWarningResponseEarlyWarnings) SetTimeBegin(v string) *DescribeEarlyWarningResponseEarlyWarnings {
	s.TimeBegin = &v
	return s
}

func (s *DescribeEarlyWarningResponseEarlyWarnings) SetTimeEnd(v string) *DescribeEarlyWarningResponseEarlyWarnings {
	s.TimeEnd = &v
	return s
}

func (s *DescribeEarlyWarningResponseEarlyWarnings) SetChannel(v string) *DescribeEarlyWarningResponseEarlyWarnings {
	s.Channel = &v
	return s
}

type DescribeConfigNameRequest struct {
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeConfigNameRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigNameRequest) SetSourceIp(v string) *DescribeConfigNameRequest {
	s.SourceIp = &v
	return s
}

type DescribeConfigNameResponse struct {
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	HasConfig   *bool                                    `json:"HasConfig,omitempty" xml:"HasConfig,omitempty" require:"true"`
	BizCode     *string                                  `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	ConfigNames []*DescribeConfigNameResponseConfigNames `json:"ConfigNames,omitempty" xml:"ConfigNames,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeConfigNameResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigNameResponse) SetRequestId(v string) *DescribeConfigNameResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeConfigNameResponse) SetHasConfig(v bool) *DescribeConfigNameResponse {
	s.HasConfig = &v
	return s
}

func (s *DescribeConfigNameResponse) SetBizCode(v string) *DescribeConfigNameResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeConfigNameResponse) SetConfigNames(v []*DescribeConfigNameResponseConfigNames) *DescribeConfigNameResponse {
	s.ConfigNames = v
	return s
}

type DescribeConfigNameResponseConfigNames struct {
	AliUid     *string `json:"AliUid,omitempty" xml:"AliUid,omitempty" require:"true"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty" require:"true"`
	RefExtId   *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty" require:"true"`
}

func (s DescribeConfigNameResponseConfigNames) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigNameResponseConfigNames) GoString() string {
	return s.String()
}

func (s *DescribeConfigNameResponseConfigNames) SetAliUid(v string) *DescribeConfigNameResponseConfigNames {
	s.AliUid = &v
	return s
}

func (s *DescribeConfigNameResponseConfigNames) SetConfigName(v string) *DescribeConfigNameResponseConfigNames {
	s.ConfigName = &v
	return s
}

func (s *DescribeConfigNameResponseConfigNames) SetRefExtId(v string) *DescribeConfigNameResponseConfigNames {
	s.RefExtId = &v
	return s
}

type DescribeCaptchaRiskRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	RefExtId   *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty"`
}

func (s DescribeCaptchaRiskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaRiskRequest) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaRiskRequest) SetSourceIp(v string) *DescribeCaptchaRiskRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCaptchaRiskRequest) SetConfigName(v string) *DescribeCaptchaRiskRequest {
	s.ConfigName = &v
	return s
}

func (s *DescribeCaptchaRiskRequest) SetTime(v string) *DescribeCaptchaRiskRequest {
	s.Time = &v
	return s
}

func (s *DescribeCaptchaRiskRequest) SetRefExtId(v string) *DescribeCaptchaRiskRequest {
	s.RefExtId = &v
	return s
}

type DescribeCaptchaRiskResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode        *string `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	NumOfThisMonth *int    `json:"NumOfThisMonth,omitempty" xml:"NumOfThisMonth,omitempty" require:"true"`
	NumOfLastMonth *int    `json:"NumOfLastMonth,omitempty" xml:"NumOfLastMonth,omitempty" require:"true"`
	RiskLevel      *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" require:"true"`
}

func (s DescribeCaptchaRiskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaRiskResponse) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaRiskResponse) SetRequestId(v string) *DescribeCaptchaRiskResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCaptchaRiskResponse) SetBizCode(v string) *DescribeCaptchaRiskResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeCaptchaRiskResponse) SetNumOfThisMonth(v int) *DescribeCaptchaRiskResponse {
	s.NumOfThisMonth = &v
	return s
}

func (s *DescribeCaptchaRiskResponse) SetNumOfLastMonth(v int) *DescribeCaptchaRiskResponse {
	s.NumOfLastMonth = &v
	return s
}

func (s *DescribeCaptchaRiskResponse) SetRiskLevel(v string) *DescribeCaptchaRiskResponse {
	s.RiskLevel = &v
	return s
}

type DescribeCaptchaMinRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	RefExtId   *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty"`
}

func (s DescribeCaptchaMinRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaMinRequest) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaMinRequest) SetSourceIp(v string) *DescribeCaptchaMinRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCaptchaMinRequest) SetConfigName(v string) *DescribeCaptchaMinRequest {
	s.ConfigName = &v
	return s
}

func (s *DescribeCaptchaMinRequest) SetType(v string) *DescribeCaptchaMinRequest {
	s.Type = &v
	return s
}

func (s *DescribeCaptchaMinRequest) SetTime(v string) *DescribeCaptchaMinRequest {
	s.Time = &v
	return s
}

func (s *DescribeCaptchaMinRequest) SetRefExtId(v string) *DescribeCaptchaMinRequest {
	s.RefExtId = &v
	return s
}

type DescribeCaptchaMinResponse struct {
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode     *string                                  `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	HasData     *bool                                    `json:"HasData,omitempty" xml:"HasData,omitempty" require:"true"`
	CaptchaMins []*DescribeCaptchaMinResponseCaptchaMins `json:"CaptchaMins,omitempty" xml:"CaptchaMins,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCaptchaMinResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaMinResponse) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaMinResponse) SetRequestId(v string) *DescribeCaptchaMinResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCaptchaMinResponse) SetBizCode(v string) *DescribeCaptchaMinResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeCaptchaMinResponse) SetHasData(v bool) *DescribeCaptchaMinResponse {
	s.HasData = &v
	return s
}

func (s *DescribeCaptchaMinResponse) SetCaptchaMins(v []*DescribeCaptchaMinResponseCaptchaMins) *DescribeCaptchaMinResponse {
	s.CaptchaMins = v
	return s
}

type DescribeCaptchaMinResponseCaptchaMins struct {
	Time         *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	Pass         *string `json:"Pass,omitempty" xml:"Pass,omitempty" require:"true"`
	Interception *string `json:"Interception,omitempty" xml:"Interception,omitempty" require:"true"`
}

func (s DescribeCaptchaMinResponseCaptchaMins) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaMinResponseCaptchaMins) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaMinResponseCaptchaMins) SetTime(v string) *DescribeCaptchaMinResponseCaptchaMins {
	s.Time = &v
	return s
}

func (s *DescribeCaptchaMinResponseCaptchaMins) SetPass(v string) *DescribeCaptchaMinResponseCaptchaMins {
	s.Pass = &v
	return s
}

func (s *DescribeCaptchaMinResponseCaptchaMins) SetInterception(v string) *DescribeCaptchaMinResponseCaptchaMins {
	s.Interception = &v
	return s
}

type DescribeCaptchaIpCityRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	RefExtId   *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty"`
}

func (s DescribeCaptchaIpCityRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaIpCityRequest) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaIpCityRequest) SetSourceIp(v string) *DescribeCaptchaIpCityRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCaptchaIpCityRequest) SetConfigName(v string) *DescribeCaptchaIpCityRequest {
	s.ConfigName = &v
	return s
}

func (s *DescribeCaptchaIpCityRequest) SetType(v string) *DescribeCaptchaIpCityRequest {
	s.Type = &v
	return s
}

func (s *DescribeCaptchaIpCityRequest) SetTime(v string) *DescribeCaptchaIpCityRequest {
	s.Time = &v
	return s
}

func (s *DescribeCaptchaIpCityRequest) SetRefExtId(v string) *DescribeCaptchaIpCityRequest {
	s.RefExtId = &v
	return s
}

type DescribeCaptchaIpCityResponse struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode       *string                                       `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	HasData       *bool                                         `json:"HasData,omitempty" xml:"HasData,omitempty" require:"true"`
	CaptchaCities []*DescribeCaptchaIpCityResponseCaptchaCities `json:"CaptchaCities,omitempty" xml:"CaptchaCities,omitempty" require:"true" type:"Repeated"`
	CaptchaIps    []*DescribeCaptchaIpCityResponseCaptchaIps    `json:"CaptchaIps,omitempty" xml:"CaptchaIps,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeCaptchaIpCityResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaIpCityResponse) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaIpCityResponse) SetRequestId(v string) *DescribeCaptchaIpCityResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCaptchaIpCityResponse) SetBizCode(v string) *DescribeCaptchaIpCityResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeCaptchaIpCityResponse) SetHasData(v bool) *DescribeCaptchaIpCityResponse {
	s.HasData = &v
	return s
}

func (s *DescribeCaptchaIpCityResponse) SetCaptchaCities(v []*DescribeCaptchaIpCityResponseCaptchaCities) *DescribeCaptchaIpCityResponse {
	s.CaptchaCities = v
	return s
}

func (s *DescribeCaptchaIpCityResponse) SetCaptchaIps(v []*DescribeCaptchaIpCityResponseCaptchaIps) *DescribeCaptchaIpCityResponse {
	s.CaptchaIps = v
	return s
}

type DescribeCaptchaIpCityResponseCaptchaCities struct {
	Location *string `json:"Location,omitempty" xml:"Location,omitempty" require:"true"`
	Lat      *string `json:"Lat,omitempty" xml:"Lat,omitempty" require:"true"`
	Lng      *string `json:"Lng,omitempty" xml:"Lng,omitempty" require:"true"`
	Pv       *int    `json:"Pv,omitempty" xml:"Pv,omitempty" require:"true"`
}

func (s DescribeCaptchaIpCityResponseCaptchaCities) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaIpCityResponseCaptchaCities) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaIpCityResponseCaptchaCities) SetLocation(v string) *DescribeCaptchaIpCityResponseCaptchaCities {
	s.Location = &v
	return s
}

func (s *DescribeCaptchaIpCityResponseCaptchaCities) SetLat(v string) *DescribeCaptchaIpCityResponseCaptchaCities {
	s.Lat = &v
	return s
}

func (s *DescribeCaptchaIpCityResponseCaptchaCities) SetLng(v string) *DescribeCaptchaIpCityResponseCaptchaCities {
	s.Lng = &v
	return s
}

func (s *DescribeCaptchaIpCityResponseCaptchaCities) SetPv(v int) *DescribeCaptchaIpCityResponseCaptchaCities {
	s.Pv = &v
	return s
}

type DescribeCaptchaIpCityResponseCaptchaIps struct {
	Ip    *string `json:"Ip,omitempty" xml:"Ip,omitempty" require:"true"`
	Value *int    `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeCaptchaIpCityResponseCaptchaIps) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaIpCityResponseCaptchaIps) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaIpCityResponseCaptchaIps) SetIp(v string) *DescribeCaptchaIpCityResponseCaptchaIps {
	s.Ip = &v
	return s
}

func (s *DescribeCaptchaIpCityResponseCaptchaIps) SetValue(v int) *DescribeCaptchaIpCityResponseCaptchaIps {
	s.Value = &v
	return s
}

type DescribeCaptchaDayRequest struct {
	SourceIp   *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	Time       *string `json:"Time,omitempty" xml:"Time,omitempty" require:"true"`
	RefExtId   *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty"`
}

func (s DescribeCaptchaDayRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaDayRequest) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaDayRequest) SetSourceIp(v string) *DescribeCaptchaDayRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCaptchaDayRequest) SetConfigName(v string) *DescribeCaptchaDayRequest {
	s.ConfigName = &v
	return s
}

func (s *DescribeCaptchaDayRequest) SetType(v string) *DescribeCaptchaDayRequest {
	s.Type = &v
	return s
}

func (s *DescribeCaptchaDayRequest) SetTime(v string) *DescribeCaptchaDayRequest {
	s.Time = &v
	return s
}

func (s *DescribeCaptchaDayRequest) SetRefExtId(v string) *DescribeCaptchaDayRequest {
	s.RefExtId = &v
	return s
}

type DescribeCaptchaDayResponse struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode    *string                               `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	HasData    *bool                                 `json:"HasData,omitempty" xml:"HasData,omitempty" require:"true"`
	CaptchaDay *DescribeCaptchaDayResponseCaptchaDay `json:"CaptchaDay,omitempty" xml:"CaptchaDay,omitempty" require:"true" type:"Struct"`
}

func (s DescribeCaptchaDayResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaDayResponse) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaDayResponse) SetRequestId(v string) *DescribeCaptchaDayResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeCaptchaDayResponse) SetBizCode(v string) *DescribeCaptchaDayResponse {
	s.BizCode = &v
	return s
}

func (s *DescribeCaptchaDayResponse) SetHasData(v bool) *DescribeCaptchaDayResponse {
	s.HasData = &v
	return s
}

func (s *DescribeCaptchaDayResponse) SetCaptchaDay(v *DescribeCaptchaDayResponseCaptchaDay) *DescribeCaptchaDayResponse {
	s.CaptchaDay = v
	return s
}

type DescribeCaptchaDayResponseCaptchaDay struct {
	Init                        *int `json:"Init,omitempty" xml:"Init,omitempty" require:"true"`
	AskForVerify                *int `json:"AskForVerify,omitempty" xml:"AskForVerify,omitempty" require:"true"`
	DirecetStrategyInterception *int `json:"DirecetStrategyInterception,omitempty" xml:"DirecetStrategyInterception,omitempty" require:"true"`
	TwiceVerify                 *int `json:"TwiceVerify,omitempty" xml:"TwiceVerify,omitempty" require:"true"`
	Pass                        *int `json:"Pass,omitempty" xml:"Pass,omitempty" require:"true"`
	CheckTested                 *int `json:"CheckTested,omitempty" xml:"CheckTested,omitempty" require:"true"`
	UncheckTested               *int `json:"UncheckTested,omitempty" xml:"UncheckTested,omitempty" require:"true"`
	LegalSign                   *int `json:"LegalSign,omitempty" xml:"LegalSign,omitempty" require:"true"`
	MaliciousFlow               *int `json:"MaliciousFlow,omitempty" xml:"MaliciousFlow,omitempty" require:"true"`
}

func (s DescribeCaptchaDayResponseCaptchaDay) String() string {
	return tea.Prettify(s)
}

func (s DescribeCaptchaDayResponseCaptchaDay) GoString() string {
	return s.String()
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetInit(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.Init = &v
	return s
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetAskForVerify(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.AskForVerify = &v
	return s
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetDirecetStrategyInterception(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.DirecetStrategyInterception = &v
	return s
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetTwiceVerify(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.TwiceVerify = &v
	return s
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetPass(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.Pass = &v
	return s
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetCheckTested(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.CheckTested = &v
	return s
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetUncheckTested(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.UncheckTested = &v
	return s
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetLegalSign(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.LegalSign = &v
	return s
}

func (s *DescribeCaptchaDayResponseCaptchaDay) SetMaliciousFlow(v int) *DescribeCaptchaDayResponseCaptchaDay {
	s.MaliciousFlow = &v
	return s
}

type CreateConfigurationRequest struct {
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ConfigurationName   *string `json:"ConfigurationName,omitempty" xml:"ConfigurationName,omitempty" require:"true"`
	ApplyType           *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty" require:"true"`
	Scene               *string `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true"`
	MaxPV               *string `json:"MaxPV,omitempty" xml:"MaxPV,omitempty" require:"true"`
	ConfigurationMethod *string `json:"ConfigurationMethod,omitempty" xml:"ConfigurationMethod,omitempty" require:"true"`
}

func (s CreateConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigurationRequest) SetSourceIp(v string) *CreateConfigurationRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateConfigurationRequest) SetConfigurationName(v string) *CreateConfigurationRequest {
	s.ConfigurationName = &v
	return s
}

func (s *CreateConfigurationRequest) SetApplyType(v string) *CreateConfigurationRequest {
	s.ApplyType = &v
	return s
}

func (s *CreateConfigurationRequest) SetScene(v string) *CreateConfigurationRequest {
	s.Scene = &v
	return s
}

func (s *CreateConfigurationRequest) SetMaxPV(v string) *CreateConfigurationRequest {
	s.MaxPV = &v
	return s
}

func (s *CreateConfigurationRequest) SetConfigurationMethod(v string) *CreateConfigurationRequest {
	s.ConfigurationMethod = &v
	return s
}

type CreateConfigurationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode   *string `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	RefExtId  *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty" require:"true"`
}

func (s CreateConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigurationResponse) SetRequestId(v string) *CreateConfigurationResponse {
	s.RequestId = &v
	return s
}

func (s *CreateConfigurationResponse) SetBizCode(v string) *CreateConfigurationResponse {
	s.BizCode = &v
	return s
}

func (s *CreateConfigurationResponse) SetRefExtId(v string) *CreateConfigurationResponse {
	s.RefExtId = &v
	return s
}

type ConfigurationStyleRequest struct {
	SourceIp            *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ApplyType           *string `json:"ApplyType,omitempty" xml:"ApplyType,omitempty" require:"true"`
	Scene               *string `json:"Scene,omitempty" xml:"Scene,omitempty" require:"true"`
	ConfigurationMethod *string `json:"ConfigurationMethod,omitempty" xml:"ConfigurationMethod,omitempty" require:"true"`
	RefExtId            *string `json:"RefExtId,omitempty" xml:"RefExtId,omitempty"`
}

func (s ConfigurationStyleRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationStyleRequest) GoString() string {
	return s.String()
}

func (s *ConfigurationStyleRequest) SetSourceIp(v string) *ConfigurationStyleRequest {
	s.SourceIp = &v
	return s
}

func (s *ConfigurationStyleRequest) SetApplyType(v string) *ConfigurationStyleRequest {
	s.ApplyType = &v
	return s
}

func (s *ConfigurationStyleRequest) SetScene(v string) *ConfigurationStyleRequest {
	s.Scene = &v
	return s
}

func (s *ConfigurationStyleRequest) SetConfigurationMethod(v string) *ConfigurationStyleRequest {
	s.ConfigurationMethod = &v
	return s
}

func (s *ConfigurationStyleRequest) SetRefExtId(v string) *ConfigurationStyleRequest {
	s.RefExtId = &v
	return s
}

type ConfigurationStyleResponse struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode   *string                             `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
	CodeData  *ConfigurationStyleResponseCodeData `json:"CodeData,omitempty" xml:"CodeData,omitempty" require:"true" type:"Struct"`
}

func (s ConfigurationStyleResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationStyleResponse) GoString() string {
	return s.String()
}

func (s *ConfigurationStyleResponse) SetRequestId(v string) *ConfigurationStyleResponse {
	s.RequestId = &v
	return s
}

func (s *ConfigurationStyleResponse) SetBizCode(v string) *ConfigurationStyleResponse {
	s.BizCode = &v
	return s
}

func (s *ConfigurationStyleResponse) SetCodeData(v *ConfigurationStyleResponseCodeData) *ConfigurationStyleResponse {
	s.CodeData = v
	return s
}

type ConfigurationStyleResponseCodeData struct {
	Html      *string `json:"Html,omitempty" xml:"Html,omitempty" require:"true"`
	Net       *string `json:"Net,omitempty" xml:"Net,omitempty" require:"true"`
	Php       *string `json:"Php,omitempty" xml:"Php,omitempty" require:"true"`
	Python    *string `json:"Python,omitempty" xml:"Python,omitempty" require:"true"`
	Java      *string `json:"Java,omitempty" xml:"Java,omitempty" require:"true"`
	NodeJs    *string `json:"NodeJs,omitempty" xml:"NodeJs,omitempty" require:"true"`
	NetUrl    *string `json:"NetUrl,omitempty" xml:"NetUrl,omitempty" require:"true"`
	PhpUrl    *string `json:"PhpUrl,omitempty" xml:"PhpUrl,omitempty" require:"true"`
	PythonUrl *string `json:"PythonUrl,omitempty" xml:"PythonUrl,omitempty" require:"true"`
	JavaUrl   *string `json:"JavaUrl,omitempty" xml:"JavaUrl,omitempty" require:"true"`
	NodeJsUrl *string `json:"NodeJsUrl,omitempty" xml:"NodeJsUrl,omitempty" require:"true"`
}

func (s ConfigurationStyleResponseCodeData) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationStyleResponseCodeData) GoString() string {
	return s.String()
}

func (s *ConfigurationStyleResponseCodeData) SetHtml(v string) *ConfigurationStyleResponseCodeData {
	s.Html = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetNet(v string) *ConfigurationStyleResponseCodeData {
	s.Net = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetPhp(v string) *ConfigurationStyleResponseCodeData {
	s.Php = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetPython(v string) *ConfigurationStyleResponseCodeData {
	s.Python = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetJava(v string) *ConfigurationStyleResponseCodeData {
	s.Java = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetNodeJs(v string) *ConfigurationStyleResponseCodeData {
	s.NodeJs = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetNetUrl(v string) *ConfigurationStyleResponseCodeData {
	s.NetUrl = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetPhpUrl(v string) *ConfigurationStyleResponseCodeData {
	s.PhpUrl = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetPythonUrl(v string) *ConfigurationStyleResponseCodeData {
	s.PythonUrl = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetJavaUrl(v string) *ConfigurationStyleResponseCodeData {
	s.JavaUrl = &v
	return s
}

func (s *ConfigurationStyleResponseCodeData) SetNodeJsUrl(v string) *ConfigurationStyleResponseCodeData {
	s.NodeJsUrl = &v
	return s
}

type AuthenticateSigRequest struct {
	SourceIp  *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty" require:"true"`
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Sig       *string `json:"Sig,omitempty" xml:"Sig,omitempty"`
	Token     *string `json:"Token,omitempty" xml:"Token,omitempty"`
	Scene     *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	RemoteIp  *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
}

func (s AuthenticateSigRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthenticateSigRequest) GoString() string {
	return s.String()
}

func (s *AuthenticateSigRequest) SetSourceIp(v string) *AuthenticateSigRequest {
	s.SourceIp = &v
	return s
}

func (s *AuthenticateSigRequest) SetSessionId(v string) *AuthenticateSigRequest {
	s.SessionId = &v
	return s
}

func (s *AuthenticateSigRequest) SetAppKey(v string) *AuthenticateSigRequest {
	s.AppKey = &v
	return s
}

func (s *AuthenticateSigRequest) SetSig(v string) *AuthenticateSigRequest {
	s.Sig = &v
	return s
}

func (s *AuthenticateSigRequest) SetToken(v string) *AuthenticateSigRequest {
	s.Token = &v
	return s
}

func (s *AuthenticateSigRequest) SetScene(v string) *AuthenticateSigRequest {
	s.Scene = &v
	return s
}

func (s *AuthenticateSigRequest) SetRemoteIp(v string) *AuthenticateSigRequest {
	s.RemoteIp = &v
	return s
}

type AuthenticateSigResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Code      *int    `json:"Code,omitempty" xml:"Code,omitempty" require:"true"`
	Msg       *string `json:"Msg,omitempty" xml:"Msg,omitempty" require:"true"`
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty" require:"true"`
	Detail    *string `json:"Detail,omitempty" xml:"Detail,omitempty" require:"true"`
}

func (s AuthenticateSigResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthenticateSigResponse) GoString() string {
	return s.String()
}

func (s *AuthenticateSigResponse) SetRequestId(v string) *AuthenticateSigResponse {
	s.RequestId = &v
	return s
}

func (s *AuthenticateSigResponse) SetCode(v int) *AuthenticateSigResponse {
	s.Code = &v
	return s
}

func (s *AuthenticateSigResponse) SetMsg(v string) *AuthenticateSigResponse {
	s.Msg = &v
	return s
}

func (s *AuthenticateSigResponse) SetRiskLevel(v string) *AuthenticateSigResponse {
	s.RiskLevel = &v
	return s
}

func (s *AuthenticateSigResponse) SetDetail(v string) *AuthenticateSigResponse {
	s.Detail = &v
	return s
}

type AnalyzeNvcRequest struct {
	SourceIp     *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	ScoreJsonStr *string `json:"ScoreJsonStr,omitempty" xml:"ScoreJsonStr,omitempty" require:"true"`
	Data         *string `json:"Data,omitempty" xml:"Data,omitempty" require:"true"`
}

func (s AnalyzeNvcRequest) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeNvcRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeNvcRequest) SetSourceIp(v string) *AnalyzeNvcRequest {
	s.SourceIp = &v
	return s
}

func (s *AnalyzeNvcRequest) SetScoreJsonStr(v string) *AnalyzeNvcRequest {
	s.ScoreJsonStr = &v
	return s
}

func (s *AnalyzeNvcRequest) SetData(v string) *AnalyzeNvcRequest {
	s.Data = &v
	return s
}

type AnalyzeNvcResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BizCode   *string `json:"BizCode,omitempty" xml:"BizCode,omitempty" require:"true"`
}

func (s AnalyzeNvcResponse) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeNvcResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeNvcResponse) SetRequestId(v string) *AnalyzeNvcResponse {
	s.RequestId = &v
	return s
}

func (s *AnalyzeNvcResponse) SetBizCode(v string) *AnalyzeNvcResponse {
	s.BizCode = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("afs.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("afs.aliyuncs.com"),
		"ap-south-1":                  tea.String("afs.aliyuncs.com"),
		"ap-southeast-1":              tea.String("afs.aliyuncs.com"),
		"ap-southeast-2":              tea.String("afs.aliyuncs.com"),
		"ap-southeast-3":              tea.String("afs.aliyuncs.com"),
		"ap-southeast-5":              tea.String("afs.aliyuncs.com"),
		"cn-beijing":                  tea.String("afs.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("afs.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("afs.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("afs.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("afs.aliyuncs.com"),
		"cn-chengdu":                  tea.String("afs.aliyuncs.com"),
		"cn-edge-1":                   tea.String("afs.aliyuncs.com"),
		"cn-fujian":                   tea.String("afs.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("afs.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("afs.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("afs.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("afs.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("afs.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("afs.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("afs.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("afs.aliyuncs.com"),
		"cn-hongkong":                 tea.String("afs.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("afs.aliyuncs.com"),
		"cn-huhehaote":                tea.String("afs.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("afs.aliyuncs.com"),
		"cn-qingdao":                  tea.String("afs.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("afs.aliyuncs.com"),
		"cn-shanghai":                 tea.String("afs.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("afs.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("afs.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("afs.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("afs.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("afs.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("afs.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("afs.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("afs.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("afs.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("afs.aliyuncs.com"),
		"cn-wuhan":                    tea.String("afs.aliyuncs.com"),
		"cn-yushanfang":               tea.String("afs.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("afs.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("afs.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("afs.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("afs.aliyuncs.com"),
		"eu-central-1":                tea.String("afs.aliyuncs.com"),
		"eu-west-1":                   tea.String("afs.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("afs.aliyuncs.com"),
		"me-east-1":                   tea.String("afs.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("afs.aliyuncs.com"),
		"us-east-1":                   tea.String("afs.aliyuncs.com"),
		"us-west-1":                   tea.String("afs.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("afs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) DescribeAfsTotalConfDataWithOptions(request *DescribeAfsTotalConfDataRequest, runtime *util.RuntimeOptions) (_result *DescribeAfsTotalConfDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAfsTotalConfDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAfsTotalConfData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAfsTotalConfData(request *DescribeAfsTotalConfDataRequest) (_result *DescribeAfsTotalConfDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAfsTotalConfDataResponse{}
	_body, _err := client.DescribeAfsTotalConfDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAfsOneConfDataWithOptions(request *DescribeAfsOneConfDataRequest, runtime *util.RuntimeOptions) (_result *DescribeAfsOneConfDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAfsOneConfDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAfsOneConfData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAfsOneConfData(request *DescribeAfsOneConfDataRequest) (_result *DescribeAfsOneConfDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAfsOneConfDataResponse{}
	_body, _err := client.DescribeAfsOneConfDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAfsVerifySigDataWithOptions(request *DescribeAfsVerifySigDataRequest, runtime *util.RuntimeOptions) (_result *DescribeAfsVerifySigDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAfsVerifySigDataResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAfsVerifySigData"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAfsVerifySigData(request *DescribeAfsVerifySigDataRequest) (_result *DescribeAfsVerifySigDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAfsVerifySigDataResponse{}
	_body, _err := client.DescribeAfsVerifySigDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAfsConfigNameWithOptions(request *DescribeAfsConfigNameRequest, runtime *util.RuntimeOptions) (_result *DescribeAfsConfigNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAfsConfigNameResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAfsConfigName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAfsConfigName(request *DescribeAfsConfigNameRequest) (_result *DescribeAfsConfigNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAfsConfigNameResponse{}
	_body, _err := client.DescribeAfsConfigNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigNameWithOptions(request *UpdateConfigNameRequest, runtime *util.RuntimeOptions) (_result *UpdateConfigNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateConfigNameResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateConfigName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigName(request *UpdateConfigNameRequest) (_result *UpdateConfigNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConfigNameResponse{}
	_body, _err := client.UpdateConfigNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCaptchaOrderWithOptions(request *DescribeCaptchaOrderRequest, runtime *util.RuntimeOptions) (_result *DescribeCaptchaOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCaptchaOrderResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCaptchaOrder"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCaptchaOrder(request *DescribeCaptchaOrderRequest) (_result *DescribeCaptchaOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCaptchaOrderResponse{}
	_body, _err := client.DescribeCaptchaOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOrderInfoWithOptions(request *DescribeOrderInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeOrderInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeOrderInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeOrderInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOrderInfo(request *DescribeOrderInfoRequest) (_result *DescribeOrderInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOrderInfoResponse{}
	_body, _err := client.DescribeOrderInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetEarlyWarningWithOptions(request *SetEarlyWarningRequest, runtime *util.RuntimeOptions) (_result *SetEarlyWarningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetEarlyWarningResponse{}
	_body, _err := client.DoRequest(tea.String("SetEarlyWarning"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetEarlyWarning(request *SetEarlyWarningRequest) (_result *SetEarlyWarningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetEarlyWarningResponse{}
	_body, _err := client.SetEarlyWarningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePersonMachineListWithOptions(request *DescribePersonMachineListRequest, runtime *util.RuntimeOptions) (_result *DescribePersonMachineListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribePersonMachineListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribePersonMachineList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePersonMachineList(request *DescribePersonMachineListRequest) (_result *DescribePersonMachineListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePersonMachineListResponse{}
	_body, _err := client.DescribePersonMachineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEarlyWarningWithOptions(request *DescribeEarlyWarningRequest, runtime *util.RuntimeOptions) (_result *DescribeEarlyWarningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeEarlyWarningResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeEarlyWarning"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEarlyWarning(request *DescribeEarlyWarningRequest) (_result *DescribeEarlyWarningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEarlyWarningResponse{}
	_body, _err := client.DescribeEarlyWarningWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigNameWithOptions(request *DescribeConfigNameRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeConfigNameResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeConfigName"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigName(request *DescribeConfigNameRequest) (_result *DescribeConfigNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigNameResponse{}
	_body, _err := client.DescribeConfigNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCaptchaRiskWithOptions(request *DescribeCaptchaRiskRequest, runtime *util.RuntimeOptions) (_result *DescribeCaptchaRiskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCaptchaRiskResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCaptchaRisk"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCaptchaRisk(request *DescribeCaptchaRiskRequest) (_result *DescribeCaptchaRiskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCaptchaRiskResponse{}
	_body, _err := client.DescribeCaptchaRiskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCaptchaMinWithOptions(request *DescribeCaptchaMinRequest, runtime *util.RuntimeOptions) (_result *DescribeCaptchaMinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCaptchaMinResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCaptchaMin"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCaptchaMin(request *DescribeCaptchaMinRequest) (_result *DescribeCaptchaMinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCaptchaMinResponse{}
	_body, _err := client.DescribeCaptchaMinWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCaptchaIpCityWithOptions(request *DescribeCaptchaIpCityRequest, runtime *util.RuntimeOptions) (_result *DescribeCaptchaIpCityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCaptchaIpCityResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCaptchaIpCity"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCaptchaIpCity(request *DescribeCaptchaIpCityRequest) (_result *DescribeCaptchaIpCityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCaptchaIpCityResponse{}
	_body, _err := client.DescribeCaptchaIpCityWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCaptchaDayWithOptions(request *DescribeCaptchaDayRequest, runtime *util.RuntimeOptions) (_result *DescribeCaptchaDayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeCaptchaDayResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeCaptchaDay"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCaptchaDay(request *DescribeCaptchaDayRequest) (_result *DescribeCaptchaDayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCaptchaDayResponse{}
	_body, _err := client.DescribeCaptchaDayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigurationWithOptions(request *CreateConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateConfigurationResponse{}
	_body, _err := client.DoRequest(tea.String("CreateConfiguration"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfiguration(request *CreateConfigurationRequest) (_result *CreateConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConfigurationResponse{}
	_body, _err := client.CreateConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigurationStyleWithOptions(request *ConfigurationStyleRequest, runtime *util.RuntimeOptions) (_result *ConfigurationStyleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfigurationStyleResponse{}
	_body, _err := client.DoRequest(tea.String("ConfigurationStyle"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigurationStyle(request *ConfigurationStyleRequest) (_result *ConfigurationStyleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigurationStyleResponse{}
	_body, _err := client.ConfigurationStyleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AuthenticateSigWithOptions(request *AuthenticateSigRequest, runtime *util.RuntimeOptions) (_result *AuthenticateSigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AuthenticateSigResponse{}
	_body, _err := client.DoRequest(tea.String("AuthenticateSig"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AuthenticateSig(request *AuthenticateSigRequest) (_result *AuthenticateSigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AuthenticateSigResponse{}
	_body, _err := client.AuthenticateSigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AnalyzeNvcWithOptions(request *AnalyzeNvcRequest, runtime *util.RuntimeOptions) (_result *AnalyzeNvcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AnalyzeNvcResponse{}
	_body, _err := client.DoRequest(tea.String("AnalyzeNvc"), tea.String("HTTPS"), tea.String("POST"), tea.String("2018-01-12"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AnalyzeNvc(request *AnalyzeNvcRequest) (_result *AnalyzeNvcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AnalyzeNvcResponse{}
	_body, _err := client.AnalyzeNvcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
